package service

import (
	"github.com/tillpaid/jira-work-log-sender/internal/model"
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
)

func ShouldHighlightTimeForWorkLog(workLog model.WorkLog, workLogTime *model.WorkLogTime, cfg *resource.Config) bool {
	if workLog.ExcludedFromSpentTimeHighlight {
		return false
	}

	thresholdHours := cfg.Highlighting.DefaultThresholdHours
	if tagThreshold, ok := cfg.Highlighting.TagSpecificThresholds[workLog.Tag]; ok {
		thresholdHours = tagThreshold
	}

	return workLogTime.Hours >= thresholdHours
}
