package import_data

import (
	"reflect"
	"testing"

	"github.com/tillpaid/jira-work-log-sender/internal/model"
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
)

func newTestConfig() *resource.Config {
	return &resource.Config{
		ForbiddenProjects: resource.ForbiddenProjects{},
		Tags:              resource.TagsConfig{Allowed: []string{}},
		Highlighting:      resource.HighlightingConfig{ExcludedIssues: []string{}},
	}
}

func Test_splitIssueNumbers(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want []string
	}{
		{
			name: "Single issue",
			raw:  "COMP-123",
			want: []string{"COMP-123"},
		},
		{
			name: "Multiple issues",
			raw:  "COMP-123,COMP-456,COMP-789",
			want: []string{"COMP-123", "COMP-456", "COMP-789"},
		},
		{
			name: "Multiple issues with spaces around commas",
			raw:  "COMP-123, COMP-456 , COMP-789",
			want: []string{"COMP-123", "COMP-456", "COMP-789"},
		},
		{
			name: "Two issues",
			raw:  "SUPPORT-1,SUPPORT-2",
			want: []string{"SUPPORT-1", "SUPPORT-2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitIssueNumbers(tt.raw)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitIssueNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildWorklogFromSection(t *testing.T) {
	tests := []struct {
		name      string
		cfg       *resource.Config
		section   []string
		number    int
		wantLen   int
		wantIssues []string
		wantTime  model.WorklogTime
		wantErr   bool
	}{
		{
			name:       "Single issue",
			cfg:        newTestConfig(),
			section:    []string{"Task | COMP-123 1h30m", "[BE]", "Some description"},
			number:     1,
			wantLen:    1,
			wantIssues: []string{"COMP-123"},
			wantTime:   model.WorklogTime{Hours: 1, Minutes: 30},
			wantErr:    false,
		},
		{
			name:       "Multiple issues comma-separated",
			cfg:        newTestConfig(),
			section:    []string{"Task | COMP-123,COMP-456,COMP-789 1h", "[BE]", "Some description"},
			number:     2,
			wantLen:    3,
			wantIssues: []string{"COMP-123", "COMP-456", "COMP-789"},
			wantTime:   model.WorklogTime{Hours: 1, Minutes: 0},
			wantErr:    false,
		},
		{
			name:       "Two issues comma-separated",
			cfg:        newTestConfig(),
			section:    []string{"Task | COMP-1,COMP-2 30m", "[BE]", "Description"},
			number:     1,
			wantLen:    2,
			wantIssues: []string{"COMP-1", "COMP-2"},
			wantTime:   model.WorklogTime{Hours: 0, Minutes: 30},
			wantErr:    false,
		},
		{
			name: "Forbidden project in comma-separated list",
			cfg: &resource.Config{
				ForbiddenProjects: resource.ForbiddenProjects{"FORBIDDEN"},
				Tags:              resource.TagsConfig{Allowed: []string{}},
				Highlighting:      resource.HighlightingConfig{ExcludedIssues: []string{}},
			},
			section: []string{"Task | COMP-123,FORBIDDEN-456 1h", "[BE]", "Description"},
			number:  1,
			wantErr: true,
		},
		{
			name:    "No description section",
			cfg:     newTestConfig(),
			section: []string{"Task | COMP-123 1h"},
			number:  1,
			wantErr: true,
		},
		{
			name:    "Invalid main information",
			cfg:     newTestConfig(),
			section: []string{"Task without pipe", "[BE]", "Description"},
			number:  1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := buildWorklogFromSection(tt.cfg, tt.section, tt.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildWorklogFromSection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}

			if len(got) != tt.wantLen {
				t.Errorf("buildWorklogFromSection() len = %d, want %d", len(got), tt.wantLen)
				return
			}

			for i, worklog := range got {
				if worklog.IssueNumber != tt.wantIssues[i] {
					t.Errorf("buildWorklogFromSection()[%d].IssueNumber = %v, want %v", i, worklog.IssueNumber, tt.wantIssues[i])
				}
				if !reflect.DeepEqual(worklog.OriginalTime, tt.wantTime) {
					t.Errorf("buildWorklogFromSection()[%d].OriginalTime = %v, want %v", i, worklog.OriginalTime, tt.wantTime)
				}
				if worklog.Number != tt.number {
					t.Errorf("buildWorklogFromSection()[%d].Number = %d, want %d", i, worklog.Number, tt.number)
				}
			}

			// All worklogs from same section must share the same description and tag
			for i := 1; i < len(got); i++ {
				if got[i].Description != got[0].Description {
					t.Errorf("buildWorklogFromSection()[%d].Description differs from [0]", i)
				}
				if got[i].Tag != got[0].Tag {
					t.Errorf("buildWorklogFromSection()[%d].Tag differs from [0]", i)
				}
				if got[i].HeaderText != got[0].HeaderText {
					t.Errorf("buildWorklogFromSection()[%d].HeaderText differs from [0]", i)
				}
			}
		})
	}
}

func Test_getMainInformation(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name        string
		args        args
		issueNumber string
		worklog     model.WorklogTime
		wantErr     bool
	}{
		{
			name:        "Valid input",
			args:        args{line: "1 - Some name | SUPPORT-123 1h30m"},
			issueNumber: "SUPPORT-123",
			worklog:     model.WorklogTime{Hours: 1, Minutes: 30},
			wantErr:     false,
		}, {
			name:        "Valid input no number",
			args:        args{line: "Some name | SUPPORT-123 1h30m"},
			issueNumber: "SUPPORT-123",
			worklog:     model.WorklogTime{Hours: 1, Minutes: 30},
			wantErr:     false,
		}, {
			name:        "Valid input no name but space",
			args:        args{line: " | SUPPORT-123 1h30m"},
			issueNumber: "SUPPORT-123",
			worklog:     model.WorklogTime{Hours: 1, Minutes: 30},
			wantErr:     false,
		}, {
			name:        "Valid input no name",
			args:        args{line: "| SUPPORT-123 1h30m"},
			issueNumber: "SUPPORT-123",
			worklog:     model.WorklogTime{Hours: 1, Minutes: 30},
			wantErr:     false,
		}, {
			name:        "Valid input no space after pipe",
			args:        args{line: "|SUPPORT-123 1h30m"},
			issueNumber: "SUPPORT-123",
			worklog:     model.WorklogTime{Hours: 1, Minutes: 30},
			wantErr:     false,
		}, {
			name:        "Valid input no time but spaces",
			args:        args{line: "1 - Some name | SUPPORT-123     "},
			issueNumber: "SUPPORT-123",
			worklog:     model.WorklogTime{Hours: 0, Minutes: 0},
			wantErr:     false,
		}, {
			name:        "Valid input no time but space",
			args:        args{line: "1 - Some name | SUPPORT-123 "},
			issueNumber: "SUPPORT-123",
			worklog:     model.WorklogTime{Hours: 0, Minutes: 0},
			wantErr:     false,
		}, {
			name:        "Valid input no time",
			args:        args{line: "1 - Some name | SUPPORT-123"},
			issueNumber: "SUPPORT-123",
			worklog:     model.WorklogTime{Hours: 0, Minutes: 0},
			wantErr:     false,
		}, {
			name:        "Invalid input space between time",
			args:        args{line: "1 - Some name | SUPPORT-123 1h 30m"},
			issueNumber: "",
			worklog:     model.WorklogTime{Hours: 0, Minutes: 0},
			wantErr:     true,
		}, {
			name:        "Valid input missing issue number",
			args:        args{line: "1 - Some name | SUPPORT-"},
			issueNumber: "SUPPORT-",
			worklog:     model.WorklogTime{Hours: 0, Minutes: 0},
			wantErr:     false,
		}, {
			name:        "Valid input issue number with only dash",
			args:        args{line: "1 - Some name | -123"},
			issueNumber: "-123",
			worklog:     model.WorklogTime{Hours: 0, Minutes: 0},
			wantErr:     false,
		}, {
			name:        "Invalid input missing issue number and time",
			args:        args{line: "1 - Some name | - "},
			issueNumber: "-",
			worklog:     model.WorklogTime{Hours: 0, Minutes: 0},
			wantErr:     false,
		}, {
			name:        "Invalid input missing issue number and time but spaces",
			args:        args{line: "1 - Some name |     -    "},
			issueNumber: "-",
			worklog:     model.WorklogTime{Hours: 0, Minutes: 0},
			wantErr:     false,
		}, {
			name:        "Invalid input only dash",
			args:        args{line: "1 - Some name | -"},
			issueNumber: "-",
			worklog:     model.WorklogTime{Hours: 0, Minutes: 0},
			wantErr:     false,
		}, {
			name:        "Invalid input missing everything after pipe",
			args:        args{line: "1 - Some name | "},
			issueNumber: "",
			worklog:     model.WorklogTime{Hours: 0, Minutes: 0},
			wantErr:     true,
		}, {
			name:        "Invalid input missing everything after pipe many spaces",
			args:        args{line: "1 - Some name |           "},
			issueNumber: "",
			worklog:     model.WorklogTime{Hours: 0, Minutes: 0},
			wantErr:     true,
		}, {
			name:        "Invalid input only name and pipe",
			args:        args{line: "1 - Some name |"},
			issueNumber: "",
			worklog:     model.WorklogTime{Hours: 0, Minutes: 0},
			wantErr:     true,
		}, {
			name:        "Valid input time without issue number",
			args:        args{line: "1 - Some name | 1h30m"},
			issueNumber: "1h30m",
			worklog:     model.WorklogTime{Hours: 0, Minutes: 0},
			wantErr:     false,
		}, {
			name:        "Invalid input no space before time",
			args:        args{line: "1 - Some name |1h30m"},
			issueNumber: "1h30m",
			worklog:     model.WorklogTime{Hours: 0, Minutes: 0},
			wantErr:     false,
		}, {
			name:        "Invalid input no pipe or issue number",
			args:        args{line: "1 - Some name"},
			issueNumber: "",
			worklog:     model.WorklogTime{Hours: 0, Minutes: 0},
			wantErr:     true,
		}, {
			name:        "Invalid input only pipe",
			args:        args{line: "|"},
			issueNumber: "",
			worklog:     model.WorklogTime{Hours: 0, Minutes: 0},
			wantErr:     true,
		}, {
			name:        "Invalid input only pipe with time",
			args:        args{line: "| 1h30m"},
			issueNumber: "1h30m",
			worklog:     model.WorklogTime{Hours: 0, Minutes: 0},
			wantErr:     false,
		}, {
			name:        "Invalid input standalone time",
			args:        args{line: "1h30m"},
			issueNumber: "",
			worklog:     model.WorklogTime{Hours: 0, Minutes: 0},
			wantErr:     true,
		}, {
			name:        "Invalid input arbitrary text",
			args:        args{line: "Some text"},
			issueNumber: "",
			worklog:     model.WorklogTime{Hours: 0, Minutes: 0},
			wantErr:     true,
		}, {
			name:        "Invalid input empty line",
			args:        args{line: ""},
			issueNumber: "",
			worklog:     model.WorklogTime{Hours: 0, Minutes: 0},
			wantErr:     true,
		}, {
			name:        "Invalid input dash and pipe with spaces",
			args:        args{line: " - | "},
			issueNumber: "",
			worklog:     model.WorklogTime{Hours: 0, Minutes: 0},
			wantErr:     true,
		}, {
			name:        "Invalid input dash and pipe",
			args:        args{line: "- |"},
			issueNumber: "",
			worklog:     model.WorklogTime{Hours: 0, Minutes: 0},
			wantErr:     true,
		}, {
			name:        "Invalid input dash and pipe without spaces",
			args:        args{line: "-|"},
			issueNumber: "",
			worklog:     model.WorklogTime{Hours: 0, Minutes: 0},
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := getMainInformation(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("getMainInformation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.issueNumber {
				t.Errorf("getMainInformation() got = %v, want %v", got, tt.issueNumber)
			}
			if !reflect.DeepEqual(got1, tt.worklog) {
				t.Errorf("getMainInformation() got1 = %v, want %v", got1, tt.worklog)
			}
		})
	}
}
