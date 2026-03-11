package page_worklogs

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

type TimeRow struct {
	Elements []*TimeRowElement
}

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

func getTimeRow(worklogs []model.Worklog, cfg *resource.Config) *TimeRow {
	total, totalModified := getTotalInMinutes(worklogs)
	left := cfg.TimeAdjustment.TargetDailyMinutes - total

	var leftColor int16 = ui.DefaultColor
	if left > cfg.TimeAdjustment.RemainingTimeThreshold || left < 0 {
		leftColor = ui.YellowOnBlack
	}

	return &TimeRow{[]*TimeRowElement{
		NewTimeRowElement(totalTimeText, total, ui.DefaultColor),
		NewTimeRowElement(leftText, left, leftColor),
		NewTimeRowElement(totalModifiedText, totalModified, ui.DefaultColor),
	}}
}

func (tr *TimeRow) GetTotalTextLen(additionalSpaces int) int {
	totalLen := 0

	for _, e := range tr.Elements {
		totalLen += len(e.Text) + additionalSpaces
	}

	return totalLen - additionalSpaces
}

func getTotalInMinutes(worklogs []model.Worklog) (int, int) {
	var total int
	var totalModified int

	for _, worklog := range worklogs {
		total += worklog.OriginalTime.GetInMinutes()
		totalModified += worklog.ModifiedTime.GetInMinutes()
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
