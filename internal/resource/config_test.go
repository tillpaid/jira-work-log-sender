package resource

import (
	"os"
	"testing"
)

func TestApplyTargetDailyTimeArgument(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    int
		wantErr bool
	}{
		{"1. No argument keeps config value", []string{"tt"}, 480, false},
		{"2. Only flags keep config value", []string{"tt", "--dev"}, 480, false},
		{"3. Hours and minutes", []string{"tt", "8h12m"}, 492, false},
		{"4. Only hours", []string{"tt", "6h"}, 360, false},
		{"5. Only minutes", []string{"tt", "90m"}, 90, false},
		{"6. Argument after flag", []string{"tt", "--dev", "7h30m"}, 450, false},
		{"7. Reversed order is invalid", []string{"tt", "12m32h"}, 480, true},
		{"8. Garbage is invalid", []string{"tt", "abc"}, 480, true},
		{"9. Zero time is invalid", []string{"tt", "0h0m"}, 480, true},
	}

	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args
			cfg := &Config{TimeAdjustment: TimeAdjustmentConfig{TargetDailyMinutes: 480}}

			err := applyTargetDailyTimeArgument(cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("applyTargetDailyTimeArgument() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if cfg.TimeAdjustment.TargetDailyMinutes != tt.want {
				t.Errorf("TargetDailyMinutes = %v, want %v", cfg.TimeAdjustment.TargetDailyMinutes, tt.want)
			}
		})
	}
}
