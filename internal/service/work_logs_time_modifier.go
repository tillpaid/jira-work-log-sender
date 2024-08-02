package service

import "github.com/tillpaid/paysera-log-time-golang/internal/model"

const (
	excluded = "TIME-505"
)

func ModifyWorkLogsTime(workLogs []model.WorkLog) []model.WorkLog {
	totalInMinutes := getTotalInMinutes(workLogs)
	totalLeft := 480 - totalInMinutes
	leftToAdd := totalLeft

	if totalLeft <= 0 {
		return workLogs
	}

	for i, workLog := range workLogs {
		if workLog.IssueNumber == excluded {
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
