package service

import (
	"os"

	"github.com/tillpaid/paysera-log-time-golang/internal/model"
)

func ModifyWorkLogsTime(workLogs []model.WorkLog) []model.WorkLog {
	totalInMinutes := getTotalInMinutes(workLogs)
	totalLeft := 480 - totalInMinutes
	leftToAdd := totalLeft

	for i, workLog := range workLogs {
		if isExcluded(workLog.IssueNumber) || totalLeft <= 0 {
			workLogs[i].ModifiedTime.Hours = workLog.OriginalTime.Hours
			workLogs[i].ModifiedTime.Minutes = workLog.OriginalTime.Minutes
			continue
		}

		inMinutes := workLog.OriginalTime.Hours*60 + workLog.OriginalTime.Minutes
		percentage := float64(inMinutes) / float64(totalInMinutes)

		toAddInMinutes := int(float64(totalLeft) * percentage)
		if i == len(workLogs)-1 {
			toAddInMinutes = leftToAdd
		}

		leftToAdd -= toAddInMinutes

		workLogs[i].ModifiedTime.Hours += (inMinutes + toAddInMinutes) / 60
		workLogs[i].ModifiedTime.Minutes += (inMinutes + toAddInMinutes) % 60

		if workLogs[i].ModifiedTime.Minutes >= 60 {
			workLogs[i].ModifiedTime.Hours++
			workLogs[i].ModifiedTime.Minutes -= 60
		}
	}

	return workLogs
}

func getTotalInMinutes(workLogs []model.WorkLog) int {
	var total int

	for _, workLog := range workLogs {
		total += workLog.OriginalTime.Hours * 60
		total += workLog.OriginalTime.Minutes
	}

	return total
}

func isExcluded(issueNumber string) bool {
	excluded := []string{"TIME-505"}

	if len(os.Args) > 1 {
		excluded = append(excluded, os.Args[1:]...)
	}

	for _, excludedIssueNumber := range excluded {
		if issueNumber == excludedIssueNumber {
			return true
		}
	}

	return false
}
