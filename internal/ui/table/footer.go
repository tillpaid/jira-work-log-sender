package table

import (
	"fmt"
	"github.com/tillpaid/paysera-log-time-golang/internal/model"
)

func GetFooter(workLogs []model.WorkLog, delimiter string) []string {
	totalInMinutes := getTotalInMinutes(workLogs)

	totalTime := calculateTotalTime(totalInMinutes)
	leftTime := calculateLeftTime(totalInMinutes)

	totalRow := fmt.Sprintf("%s | %s", totalTime, leftTime)

	return []string{delimiter, totalRow, delimiter}
}

func calculateTotalTime(totalInMinutes int) string {
	var hours, minutes int

	if totalInMinutes > 0 {
		hours = totalInMinutes / 60
		minutes = totalInMinutes % 60
	}

	return fmt.Sprintf("Total time: %dh %dm", hours, minutes)
}

func calculateLeftTime(totalInMinutes int) string {
	var hours, minutes int
	var minusSign string

	minutes = 480 - totalInMinutes
	if minutes < 0 {
		minutes = minutes * -1
		minusSign = "-"
	}

	hours += minutes / 60
	minutes = minutes % 60

	return fmt.Sprintf("Left: %s%dh %dm", minusSign, hours, minutes)
}

func getTotalInMinutes(workLogs []model.WorkLog) int {
	var total int

	for _, workLog := range workLogs {
		total += workLog.OriginalTime.Hours * 60
		total += workLog.OriginalTime.Minutes
	}

	return total
}
