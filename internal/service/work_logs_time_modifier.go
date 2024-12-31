package service

import (
	"strings"

	"github.com/tillpaid/jira-work-log-sender/internal/model"
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
)

func ModifyWorkLogsTime(workLogs []model.WorkLog, config *resource.Config) []model.WorkLog {
	totalInMinutes := getTotalInMinutes(workLogs, config, true)
	totalNotExcludedInMinutes := getTotalInMinutes(workLogs, config, false)

	if totalInMinutes == 0 {
		return workLogs
	}

	totalLeft := config.TargetTime - totalInMinutes
	leftToAdd := totalLeft
	lastUpdatedIndex := -1

	for i, workLog := range workLogs {
		workLogs[i].ModifiedTime.Hours = workLog.OriginalTime.Hours
		workLogs[i].ModifiedTime.Minutes = workLog.OriginalTime.Minutes

		if isExcluded(workLog, config) || totalLeft <= 0 {
			continue
		}

		lastUpdatedIndex = i

		hasInMinutes := workLog.OriginalTime.Hours*60 + workLog.OriginalTime.Minutes
		percentageOfTotal := float64(hasInMinutes) / float64(totalNotExcludedInMinutes)

		toAddInMinutes := int(float64(totalLeft) * percentageOfTotal)
		if i == len(workLogs)-1 {
			toAddInMinutes = leftToAdd
		}

		workLogs[i].ModifiedTime.AddMinutes(toAddInMinutes)
		leftToAdd -= toAddInMinutes
	}

	if leftToAdd > 0 && lastUpdatedIndex != -1 {
		workLogs[lastUpdatedIndex].ModifiedTime.AddMinutes(leftToAdd)
	}

	return workLogs
}

func getTotalInMinutes(workLogs []model.WorkLog, config *resource.Config, includeExcluded bool) int {
	var total int

	for _, workLog := range workLogs {
		if includeExcluded == false && isExcluded(workLog, config) {
			continue
		}

		total += workLog.OriginalTime.GetInMinutes()
	}

	return total
}

func isExcluded(workLog model.WorkLog, config *resource.Config) bool {
	if workLog.ModifyTimeDisabled || !config.TimeModification.Enabled {
		return true
	}

	issueNumber := strings.ToLower(workLog.IssueNumber)

	for _, excludedIssueNumber := range config.TimeModification.ExcludedNumbers {
		if issueNumber == strings.ToLower(excludedIssueNumber) {
			return true
		}
	}

	return false
}
