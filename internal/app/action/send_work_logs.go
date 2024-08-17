package action

import (
	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/jira"
	"github.com/tillpaid/paysera-log-time-golang/internal/model"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui/pages"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui/pages/page_send_work_logs"
)

const (
	textInProgress      = "Sending..."
	textInProgressClear = "          "
	textDone            = "Done!"
	textFailed          = "Fail!"
	textGetting         = "Calculating..."
	textGettingClear    = "              "
)

func SendLogWorks(client *jira.Client, screen *goncurses.Window, workLogs []model.WorkLog) error {
	valuesWidth := model.NewWorkLogTableWidthWithCalculations(workLogs)

	if err := pages.DrawSendWorkLogsPage(screen, workLogs, valuesWidth); err != nil {
		return err
	}

	rows := page_send_work_logs.GetBody(workLogs, valuesWidth)

	for i, workLog := range workLogs {
		if err := sendAndUpdateRow(client, screen, workLog, i+3, len(rows[i])+3); err != nil {
			return err
		}
	}

	return nil
}

func sendAndUpdateRow(client *jira.Client, screen *goncurses.Window, workLog model.WorkLog, row int, offset int) error {
	var err error

	offset, err = sendWorkLog(client, screen, workLog, row, offset)
	if err != nil {
		return err
	}

	return spentTime(client, screen, workLog, row, offset)
}

func sendWorkLog(client *jira.Client, screen *goncurses.Window, workLog model.WorkLog, row int, offset int) (int, error) {
	if err := printColored(screen, ui.YellowOnBlack, row, offset, textInProgress); err != nil {
		return offset, err
	}
	screen.Refresh()

	statusText := textDone
	var statusColor int16 = ui.GreenOnBlack

	err := client.WorkLogService.SendWorkLog(workLog)
	if err != nil {
		statusText = textFailed
		statusColor = ui.RedOnBlack
	}

	screen.MovePrint(row, offset, textInProgressClear)
	if err = printColored(screen, statusColor, row, offset, statusText); err != nil {
		return offset, err
	}

	offset += len(statusText)
	screen.MovePrint(row, offset, " | ")
	offset += 3
	screen.Refresh()

	return offset, nil
}

func spentTime(client *jira.Client, screen *goncurses.Window, workLog model.WorkLog, row int, offset int) error {
	if err := printColored(screen, ui.YellowOnBlack, row, offset, textGetting); err != nil {
		return err
	}
	screen.Refresh()

	totalTime := client.WorkLogService.GetSpentTime(workLog.IssueNumber)

	screen.MovePrint(row, offset, textGettingClear)
	screen.MovePrint(row, offset, "Total: "+totalTime)
	screen.Refresh()

	return nil
}

func printColored(screen *goncurses.Window, color int16, y int, x int, text string) error {
	if err := screen.ColorOn(color); err != nil {
		return err
	}

	screen.MovePrint(y, x, text)

	if err := screen.ColorOff(color); err != nil {
		return err
	}

	return nil
}
