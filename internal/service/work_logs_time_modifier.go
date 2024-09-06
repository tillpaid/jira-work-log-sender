package service

import (
	"os"
	"strings"

	"github.com/tillpaid/paysera-log-time-golang/internal/model"
)

func ModifyWorkLogsTime(workLogs []model.WorkLog) []model.WorkLog {
	totalInMinutes := getTotalInMinutes(workLogs, true)
	totalNotExcludedInMinutes := getTotalInMinutes(workLogs, false)

	totalLeft := 480 - totalInMinutes
	leftToAdd := totalLeft
	lastUpdatedIndex := -1

	for i, workLog := range workLogs {
		workLogs[i].ModifiedTime.Hours = workLog.OriginalTime.Hours
		workLogs[i].ModifiedTime.Minutes = workLog.OriginalTime.Minutes

		if isExcluded(workLog.IssueNumber) || totalLeft <= 0 {
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
		if includeExcluded == false && isExcluded(workLog.IssueNumber) {
			continue
		}

		total += workLog.OriginalTime.Hours * 60
		total += workLog.OriginalTime.Minutes
	}

	return total
}

func isExcluded(issueNumber string) bool {
	excluded := []string{"TIME-505"}
	issueNumber = strings.ToLower(issueNumber)

	if len(os.Args) > 1 {
		excluded = append(excluded, os.Args[1:]...)
	}

	for _, excludedIssueNumber := range excluded {
		if issueNumber == strings.ToLower(excludedIssueNumber) {
			return true
		}
	}

	return false
}
