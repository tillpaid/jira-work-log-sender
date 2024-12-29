package cache

import (
	"fmt"
	"testing"
)

func TestIssuesExistenceCache_removeHalf(t *testing.T) {
	type fields struct {
		existingIssues []string
	}

	var issues []string
	for i := 1; i <= 100; i++ {
		issues = append(issues, fmt.Sprintf("ISSUE-%d", i))
	}

	tests := []struct {
		name                string
		fields              fields
		expectedIssuesCount int
	}{
		{"1 issue", fields{issues[:1]}, 0},
		{"2 issues", fields{issues[:2]}, 1},
		{"10 issues", fields{issues[:10]}, 5},
		{"15 issues", fields{issues[:15]}, 7},
		{"20 issues", fields{issues[:20]}, 10},
		{"50 issues", fields{issues[:50]}, 25},
		{"51 issues", fields{issues[:51]}, 25},
		{"55 issues", fields{issues[:55]}, 27},
		{"73 issues", fields{issues[:73]}, 36},
		{"90 issues", fields{issues[:90]}, 45},
		{"95 issues", fields{issues[:95]}, 47},
		{"100 issues", fields{issues[:100]}, 50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var checkedExistence = make(map[string]bool, len(tt.fields.existingIssues))
			for _, issue := range tt.fields.existingIssues {
				checkedExistence[issue] = true
			}

			c := &IssuesExistenceCache{checkedExistence: checkedExistence}
			c.removeHalf()

			if tt.expectedIssuesCount != len(c.checkedExistence) {
				t.Errorf("Expected %d issues, got %d", tt.expectedIssuesCount, len(c.checkedExistence))
			}
		})
	}
}

func BenchmarkIssuesExistenceCache_removeHalfSmall(b *testing.B) {
	benchmarkRemoveHalf(b, 100)
}

func BenchmarkIssuesExistenceCache_removeHalfMedium(b *testing.B) {
	benchmarkRemoveHalf(b, 10_000)
}

func BenchmarkIssuesExistenceCache_removeHalfLarge(b *testing.B) {
	benchmarkRemoveHalf(b, 1_000_000)
}

func benchmarkRemoveHalf(b *testing.B, amountOfIssues int) {
	b.StopTimer()
	var issues = make(map[string]bool, amountOfIssues)
	for i := 1; i <= amountOfIssues; i++ {
		number := fmt.Sprintf("ISSUE-%d", i)
		issues[number] = true
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c := &IssuesExistenceCache{checkedExistence: issues}

		c.removeHalf()
	}
}
