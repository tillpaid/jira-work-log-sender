package import_data

import (
	"reflect"
	"testing"

	"github.com/tillpaid/jira-work-log-sender/internal/model"
)

func Test_getMainInformation(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name        string
		args        args
		issueNumber string
		workLog     model.WorkLogTime
		wantErr     bool
	}{
		{
			name:        "Valid input",
			args:        args{line: "1 - Some name | SUPPORT-123 1h30m"},
			issueNumber: "SUPPORT-123",
			workLog:     model.WorkLogTime{Hours: 1, Minutes: 30},
			wantErr:     false,
		}, {
			name:        "Valid input no number",
			args:        args{line: "Some name | SUPPORT-123 1h30m"},
			issueNumber: "SUPPORT-123",
			workLog:     model.WorkLogTime{Hours: 1, Minutes: 30},
			wantErr:     false,
		}, {
			name:        "Valid input no name but space",
			args:        args{line: " | SUPPORT-123 1h30m"},
			issueNumber: "SUPPORT-123",
			workLog:     model.WorkLogTime{Hours: 1, Minutes: 30},
			wantErr:     false,
		}, {
			name:        "Valid input no name",
			args:        args{line: "| SUPPORT-123 1h30m"},
			issueNumber: "SUPPORT-123",
			workLog:     model.WorkLogTime{Hours: 1, Minutes: 30},
			wantErr:     false,
		}, {
			name:        "Valid input no space after pipe",
			args:        args{line: "|SUPPORT-123 1h30m"},
			issueNumber: "SUPPORT-123",
			workLog:     model.WorkLogTime{Hours: 1, Minutes: 30},
			wantErr:     false,
		}, {
			name:        "Valid input no time but spaces",
			args:        args{line: "1 - Some name | SUPPORT-123     "},
			issueNumber: "SUPPORT-123",
			workLog:     model.WorkLogTime{Hours: 0, Minutes: 0},
			wantErr:     false,
		}, {
			name:        "Valid input no time but space",
			args:        args{line: "1 - Some name | SUPPORT-123 "},
			issueNumber: "SUPPORT-123",
			workLog:     model.WorkLogTime{Hours: 0, Minutes: 0},
			wantErr:     false,
		}, {
			name:        "Valid input no time",
			args:        args{line: "1 - Some name | SUPPORT-123"},
			issueNumber: "SUPPORT-123",
			workLog:     model.WorkLogTime{Hours: 0, Minutes: 0},
			wantErr:     false,
		}, {
			name:        "Invalid input space between time",
			args:        args{line: "1 - Some name | SUPPORT-123 1h 30m"},
			issueNumber: "",
			workLog:     model.WorkLogTime{Hours: 0, Minutes: 0},
			wantErr:     true,
		}, {
			name:        "Valid input missing issue number",
			args:        args{line: "1 - Some name | SUPPORT-"},
			issueNumber: "SUPPORT-",
			workLog:     model.WorkLogTime{Hours: 0, Minutes: 0},
			wantErr:     false,
		}, {
			name:        "Valid input issue number with only dash",
			args:        args{line: "1 - Some name | -123"},
			issueNumber: "-123",
			workLog:     model.WorkLogTime{Hours: 0, Minutes: 0},
			wantErr:     false,
		}, {
			name:        "Invalid input missing issue number and time",
			args:        args{line: "1 - Some name | - "},
			issueNumber: "-",
			workLog:     model.WorkLogTime{Hours: 0, Minutes: 0},
			wantErr:     false,
		}, {
			name:        "Invalid input missing issue number and time but spaces",
			args:        args{line: "1 - Some name |     -    "},
			issueNumber: "-",
			workLog:     model.WorkLogTime{Hours: 0, Minutes: 0},
			wantErr:     false,
		}, {
			name:        "Invalid input only dash",
			args:        args{line: "1 - Some name | -"},
			issueNumber: "-",
			workLog:     model.WorkLogTime{Hours: 0, Minutes: 0},
			wantErr:     false,
		}, {
			name:        "Invalid input missing everything after pipe",
			args:        args{line: "1 - Some name | "},
			issueNumber: "",
			workLog:     model.WorkLogTime{Hours: 0, Minutes: 0},
			wantErr:     true,
		}, {
			name:        "Invalid input missing everything after pipe many spaces",
			args:        args{line: "1 - Some name |           "},
			issueNumber: "",
			workLog:     model.WorkLogTime{Hours: 0, Minutes: 0},
			wantErr:     true,
		}, {
			name:        "Invalid input only name and pipe",
			args:        args{line: "1 - Some name |"},
			issueNumber: "",
			workLog:     model.WorkLogTime{Hours: 0, Minutes: 0},
			wantErr:     true,
		}, {
			name:        "Valid input time without issue number",
			args:        args{line: "1 - Some name | 1h30m"},
			issueNumber: "1h30m",
			workLog:     model.WorkLogTime{Hours: 0, Minutes: 0},
			wantErr:     false,
		}, {
			name:        "Invalid input no space before time",
			args:        args{line: "1 - Some name |1h30m"},
			issueNumber: "1h30m",
			workLog:     model.WorkLogTime{Hours: 0, Minutes: 0},
			wantErr:     false,
		}, {
			name:        "Invalid input no pipe or issue number",
			args:        args{line: "1 - Some name"},
			issueNumber: "",
			workLog:     model.WorkLogTime{Hours: 0, Minutes: 0},
			wantErr:     true,
		}, {
			name:        "Invalid input only pipe",
			args:        args{line: "|"},
			issueNumber: "",
			workLog:     model.WorkLogTime{Hours: 0, Minutes: 0},
			wantErr:     true,
		}, {
			name:        "Invalid input only pipe with time",
			args:        args{line: "| 1h30m"},
			issueNumber: "1h30m",
			workLog:     model.WorkLogTime{Hours: 0, Minutes: 0},
			wantErr:     false,
		}, {
			name:        "Invalid input standalone time",
			args:        args{line: "1h30m"},
			issueNumber: "",
			workLog:     model.WorkLogTime{Hours: 0, Minutes: 0},
			wantErr:     true,
		}, {
			name:        "Invalid input arbitrary text",
			args:        args{line: "Some text"},
			issueNumber: "",
			workLog:     model.WorkLogTime{Hours: 0, Minutes: 0},
			wantErr:     true,
		}, {
			name:        "Invalid input empty line",
			args:        args{line: ""},
			issueNumber: "",
			workLog:     model.WorkLogTime{Hours: 0, Minutes: 0},
			wantErr:     true,
		}, {
			name:        "Invalid input dash and pipe with spaces",
			args:        args{line: " - | "},
			issueNumber: "",
			workLog:     model.WorkLogTime{Hours: 0, Minutes: 0},
			wantErr:     true,
		}, {
			name:        "Invalid input dash and pipe",
			args:        args{line: "- |"},
			issueNumber: "",
			workLog:     model.WorkLogTime{Hours: 0, Minutes: 0},
			wantErr:     true,
		}, {
			name:        "Invalid input dash and pipe without spaces",
			args:        args{line: "-|"},
			issueNumber: "",
			workLog:     model.WorkLogTime{Hours: 0, Minutes: 0},
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
			if !reflect.DeepEqual(got1, tt.workLog) {
				t.Errorf("getMainInformation() got1 = %v, want %v", got1, tt.workLog)
			}
		})
	}
}
