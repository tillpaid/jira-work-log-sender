package table

import (
	"fmt"
	"github.com/tillpaid/paysera-log-time-golang/internal/model"
)

func GetFooter(workLogs []model.WorkLog, delimiter string) []string {
	totalTime := calculateTotalTime(workLogs)
	leftTime := calculateLeftTime(workLogs)
	totalRow := fmt.Sprintf("%s | %s", totalTime, leftTime)

	return []string{delimiter, totalRow, delimiter}
}

func calculateTotalTime(workLogs []model.WorkLog) string {
	var hours, minutes int

	for _, workLog := range workLogs {
		hours += workLog.OriginalTime.Hours
		minutes += workLog.OriginalTime.Minutes
	}

	if minutes >= 60 {
		hours += minutes / 60
		minutes = minutes % 60
	}

	return fmt.Sprintf("Total time: %dh %dm", hours, minutes)
}

func calculateLeftTime(workLogs []model.WorkLog) string {
	var totalMinutes int
	var hours, minutes int
	var minutesWasNegative bool

	for _, workLog := range workLogs {
		totalMinutes += workLog.OriginalTime.Hours * 60
		totalMinutes += workLog.OriginalTime.Minutes
	}

	minutes = 480 - totalMinutes
	if minutes < 0 {
		minutesWasNegative = true
		minutes = minutes * -1
	}

	if minutes >= 60 {
		hours += minutes / 60
		minutes = minutes % 60
	}

	minusSign := ""
	if minutesWasNegative {
		minusSign = "-"
	}

	return fmt.Sprintf("Left: %s%dh %dm", minusSign, hours, minutes)
}
