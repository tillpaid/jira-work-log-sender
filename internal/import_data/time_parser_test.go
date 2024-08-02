package import_data

import "testing"

func Test_parseTimeString(t *testing.T) {
	type args struct {
		timeString string
	}
	tests := []struct {
		name    string
		args    args
		hours   int
		minutes int
		wantErr bool
	}{
		{"1. Only minutes", args{"30m"}, 0, 30, false},
		{"2. Only hours", args{"2h"}, 2, 0, false},
		{"3. Hours and minutes", args{"2h30m"}, 2, 30, false},
		{"4. Hours and minutes reversed", args{"30m2h"}, 0, 0, true},
		{"5. Invalid time string double minutes", args{"2h30m30m"}, 0, 0, true},
		{"6. Invalid time string double hours", args{"2h30m2h"}, 0, 0, true},
		{"7. Invalid time string double hours and minutes", args{"2h30m2h30m"}, 0, 0, true},
		{"8. Invalid time string no time", args{""}, 0, 0, true},
		{"9. Invalid time string no time h", args{"h"}, 0, 0, true},
		{"10. Invalid time string no time m", args{"m"}, 0, 0, true},
		{"11. Invalid time string no time hm", args{"hm"}, 0, 0, true},
		{"12. Invalid time string no time mh", args{"mh"}, 0, 0, true},
		{"13. Invalid time string no time h2", args{"h2"}, 0, 0, true},
		{"14. Invalid time string no time m2", args{"m2"}, 0, 0, true},
		{"15. Not integer hours float", args{"2.5h"}, 0, 0, true},
		{"16. Not integer minutes float", args{"2.5m"}, 0, 0, true},
		{"17. Not integer hours and minutes float", args{"2.5h2.5m"}, 0, 0, true},
		{"18. Not integer hours string", args{"sdfh"}, 0, 0, true},
		{"19. Not integer minutes string", args{"sdfm"}, 0, 0, true},
		{"20. Not integer hours and minutes string", args{"sdfhsdfm"}, 0, 0, true},
		{"21. Negative hours", args{"-2h"}, 0, 0, true},
		{"22. Negative minutes", args{"-2m"}, 0, 0, true},
		{"23. Negative hours and minutes", args{"-2h-2m"}, 0, 0, true},
		{"24. Negative hours and minutes", args{"-2h2m"}, 0, 0, true},
		{"25. Negative hours and minutes", args{"2h-2m"}, 0, 0, true},
		{"26. Hours, minutes and seconds", args{"2h30m30s"}, 2, 30, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTimeString(tt.args.timeString)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseTimeString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			if got.Hours != tt.hours {
				t.Errorf("parseTimeString() got = %v, want %v", got.Hours, tt.hours)
			}
			if got.Minutes != tt.minutes {
				t.Errorf("parseTimeString() got1 = %v, want %v", got.Minutes, tt.minutes)
			}
		})
	}
}
