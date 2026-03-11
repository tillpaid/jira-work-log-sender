package service

import (
	"github.com/tillpaid/jira-work-log-sender/internal/model"
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
)

func ShouldHighlightTimeForWorklog(worklog model.Worklog, worklogTime *model.WorklogTime, cfg *resource.Config) bool {
	if worklog.ExcludedFromSpentTimeHighlight {
		return false
	}

	thresholdHours := cfg.Highlighting.DefaultThresholdHours
	if tagThreshold, ok := cfg.Highlighting.TagSpecificThresholds[worklog.Tag]; ok {
		thresholdHours = tagThreshold
	}

	return worklogTime.Hours >= thresholdHours
}
