package service

import (
	"strings"

	"github.com/tillpaid/jira-work-log-sender/internal/model"
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
)

func ModifyWorklogsTime(worklogs []model.Worklog, cfg *resource.Config) []model.Worklog {
	totalInMinutes := getTotalInMinutes(worklogs, cfg, true)
	totalNotExcludedInMinutes := getTotalInMinutes(worklogs, cfg, false)

	if totalInMinutes == 0 {
		return worklogs
	}

	totalLeft := cfg.TimeAdjustment.TargetDailyMinutes - totalInMinutes
	leftToAdd := totalLeft
	lastUpdatedIndex := -1

	for i, worklog := range worklogs {
		worklogs[i].ModifiedTime.Hours = worklog.OriginalTime.Hours
		worklogs[i].ModifiedTime.Minutes = worklog.OriginalTime.Minutes

		if isExcluded(worklog, cfg) || totalLeft <= 0 {
			continue
		}

		lastUpdatedIndex = i

		hasInMinutes := worklog.OriginalTime.Hours*60 + worklog.OriginalTime.Minutes
		percentageOfTotal := float64(hasInMinutes) / float64(totalNotExcludedInMinutes)

		toAddInMinutes := int(float64(totalLeft) * percentageOfTotal)
		if i == len(worklogs)-1 {
			toAddInMinutes = leftToAdd
		}

		worklogs[i].ModifiedTime.AddMinutes(toAddInMinutes)
		leftToAdd -= toAddInMinutes
	}

	if leftToAdd > 0 && lastUpdatedIndex != -1 {
		worklogs[lastUpdatedIndex].ModifiedTime.AddMinutes(leftToAdd)
	}

	return worklogs
}

func getTotalInMinutes(worklogs []model.Worklog, cfg *resource.Config, includeExcluded bool) int {
	var total int

	for _, worklog := range worklogs {
		if includeExcluded == false && isExcluded(worklog, cfg) {
			continue
		}

		total += worklog.OriginalTime.GetInMinutes()
	}

	return total
}

func isExcluded(worklog model.Worklog, cfg *resource.Config) bool {
	if worklog.ModifyTimeDisabled || !cfg.TimeAdjustment.Enabled {
		return true
	}

	issueNumber := strings.ToLower(worklog.IssueNumber)

	for _, excludedIssueNumber := range cfg.TimeAdjustment.ExcludedIssues {
		if issueNumber == strings.ToLower(excludedIssueNumber) {
			return true
		}
	}

	return false
}
