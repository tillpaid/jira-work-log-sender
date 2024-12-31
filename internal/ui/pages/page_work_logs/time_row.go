package page_work_logs

import (
	"fmt"

	"github.com/tillpaid/jira-work-log-sender/internal/model"
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
	"github.com/tillpaid/jira-work-log-sender/internal/ui"
)

const (
	totalTimeText     = "Total time"
	leftText          = "Left"
	totalModifiedText = "Total modified time"
)

type TimeRowElement struct {
	Text  string
	Color int16
}

func NewTimeRowElement(text string, minutes int, color int16) *TimeRowElement {
	return &TimeRowElement{
		Text:  fmt.Sprintf("%s: %s", text, minutesToTimeString(minutes)),
		Color: color,
	}
}

func getTimeRow(workLogs []model.WorkLog, config *resource.Config) []*TimeRowElement {
	total, totalModified := getTotalInMinutes(workLogs)
	left := config.TargetTime - total

	var leftColor int16 = ui.DefaultColor
	if left > config.RemainingTimeHighlightThreshold || left < 0 {
		leftColor = ui.YellowOnBlack
	}

	return []*TimeRowElement{
		NewTimeRowElement(totalTimeText, total, ui.DefaultColor),
		NewTimeRowElement(leftText, left, leftColor),
		NewTimeRowElement(totalModifiedText, totalModified, ui.DefaultColor),
	}
}

func getTotalInMinutes(workLogs []model.WorkLog) (int, int) {
	var total int
	var totalModified int

	for _, workLog := range workLogs {
		total += workLog.OriginalTime.GetInMinutes()
		totalModified += workLog.ModifiedTime.GetInMinutes()
	}

	return total, totalModified
}

func minutesToTimeString(timeInMinutes int) string {
	minusSign := ""

	if timeInMinutes < 0 {
		timeInMinutes *= -1
		minusSign = "-"
	}

	hours := timeInMinutes / 60
	minutes := timeInMinutes % 60

	return fmt.Sprintf("%s%dh %dm", minusSign, hours, minutes)
}
