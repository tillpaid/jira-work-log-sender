package service

import (
	"github.com/tillpaid/jira-work-log-sender/internal/model"
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
)

func ShouldHighlightTimeForWorkLog(workLog model.WorkLog, workLogTime *model.WorkLogTime, config *resource.Config) bool {
	if workLog.ExcludedFromSpentTimeHighlight {
		return false
	}

	thresholdHours := config.Highlighting.DefaultThresholdHours
	if tagThreshold, ok := config.Highlighting.TagSpecificThresholds[workLog.Tag]; ok {
		thresholdHours = tagThreshold
	}

	return workLogTime.Hours >= thresholdHours
}
