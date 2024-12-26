package service

import (
	"strings"

	"github.com/tillpaid/jira-work-log-sender/internal/model"
)

func ModifyWorkLogsTime(workLogs []model.WorkLog) []model.WorkLog {
	totalInMinutes := getTotalInMinutes(workLogs, true)
	totalNotExcludedInMinutes := getTotalInMinutes(workLogs, false)

	if totalInMinutes == 0 {
		return workLogs
	}

	totalLeft := 480 - totalInMinutes
	leftToAdd := totalLeft
	lastUpdatedIndex := -1

	for i, workLog := range workLogs {
		workLogs[i].ModifiedTime.Hours = workLog.OriginalTime.Hours
		workLogs[i].ModifiedTime.Minutes = workLog.OriginalTime.Minutes

		if isExcluded(workLog) || totalLeft <= 0 {
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

func getTotalInMinutes(workLogs []model.WorkLog, includeExcluded bool) int {
	var total int

	for _, workLog := range workLogs {
		if includeExcluded == false && isExcluded(workLog) {
			continue
		}

		total += workLog.OriginalTime.Hours * 60
		total += workLog.OriginalTime.Minutes
	}

	return total
}

func isExcluded(workLog model.WorkLog) bool {
	if workLog.ModifyTimeDisabled {
		return true
	}

	excluded := []string{"TIME-505"}
	issueNumber := strings.ToLower(workLog.IssueNumber)

	for _, excludedIssueNumber := range excluded {
		if issueNumber == strings.ToLower(excludedIssueNumber) {
			return true
		}
	}

	return false
}
