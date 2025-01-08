package service

import (
	"github.com/tillpaid/jira-work-log-sender/internal/model"
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
)

func ShouldHighlightTimeForWorkLog(workLog model.WorkLog, workLogTime *model.WorkLogTime, config *resource.Config) bool {
	if workLog.ExcludedFromSpentTimeHighlight {
		return false
	}

	highlightAfterTime := config.IssueHighlight.HighlightAfterHours
	if tagTime, ok := config.IssueHighlight.HighlightTagsAfterHours[workLog.Tag]; ok {
		highlightAfterTime = tagTime
	}

	return workLogTime.Hours >= highlightAfterTime
}
