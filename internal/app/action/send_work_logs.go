package action

import (
	"fmt"

	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/jira"
	"github.com/tillpaid/paysera-log-time-golang/internal/model"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui/pages"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui/pages/page_send_work_logs"
)

const (
	textWaiting         = "Waiting..."
	textWaitingClear    = "          "
	textInProgress      = "Sending..."
	textInProgressClear = "          "
	textDone            = "Done!"
	textFailed          = "Fail!"
	textGetting         = "Calculating..."
	textGettingClear    = "              "
)

type SendWorkLogsAction struct {
	client *jira.Client
	screen *goncurses.Window
}

func NewSendWorkLogsAction(client *jira.Client, screen *goncurses.Window) *SendWorkLogsAction {
	return &SendWorkLogsAction{client: client, screen: screen}
}

func (a *SendWorkLogsAction) Send(workLogs []model.WorkLog) error {
	for _, workLog := range workLogs {
		if workLog.OriginalTime.Hours == 0 && workLog.OriginalTime.Minutes == 0 {
			return fmt.Errorf("work log with issue number %s has no time spent", workLog.IssueNumber)
		}
	}

	_, width := a.screen.MaxYX()
	valuesWidth := model.NewWorkLogTableWidthWithCalculations(workLogs, width)

	if err := pages.DrawSendWorkLogsPage(a.screen, workLogs, valuesWidth); err != nil {
		return err
	}

	rows := page_send_work_logs.GetBody(workLogs, valuesWidth)

	for i := range workLogs {
		if err := a.setWaiting(i+3, len(rows[i])+3); err != nil {
			return err
		}
	}
	a.screen.Refresh()

	for i, workLog := range workLogs {
		if err := a.sendAndUpdateRow(workLog, i+3, len(rows[i])+3); err != nil {
			return err
		}
	}

	return nil
}

func (a *SendWorkLogsAction) setWaiting(row int, offset int) error {
	if err := pages.PrintColored(a.screen, ui.CyanOnBlack, row, offset, textWaiting); err != nil {
		return err
	}

	return nil
}

func (a *SendWorkLogsAction) sendAndUpdateRow(workLog model.WorkLog, row int, offset int) error {
	var err error

	offset, err = a.sendWorkLog(workLog, row, offset)
	if err != nil {
		return err
	}

	return a.spentTime(workLog, row, offset)
}

func (a *SendWorkLogsAction) sendWorkLog(workLog model.WorkLog, row int, offset int) (int, error) {
	a.screen.MovePrint(row, offset, textWaitingClear)
	if err := pages.PrintColored(a.screen, ui.YellowOnBlack, row, offset, textInProgress); err != nil {
		return offset, err
	}
	a.screen.Refresh()

	statusText := textDone
	var statusColor int16 = ui.GreenOnBlack

	err := a.client.WorkLogService.SendWorkLog(workLog)
	if err != nil {
		statusText = textFailed
		statusColor = ui.RedOnBlack
	}

	a.screen.MovePrint(row, offset, textInProgressClear)
	if err = pages.PrintColored(a.screen, statusColor, row, offset, statusText); err != nil {
		return offset, err
	}

	offset += len(statusText)
	a.screen.MovePrint(row, offset, " | ")
	offset += 3
	a.screen.Refresh()

	return offset, nil
}

func (a *SendWorkLogsAction) spentTime(workLog model.WorkLog, row int, offset int) error {
	if err := pages.PrintColored(a.screen, ui.YellowOnBlack, row, offset, textGetting); err != nil {
		return err
	}
	a.screen.Refresh()

	totalTime := a.client.WorkLogService.GetSpentTime(workLog.IssueNumber)

	a.screen.MovePrint(row, offset, textGettingClear)
	a.screen.MovePrint(row, offset, "Total: "+totalTime)
	a.screen.Refresh()

	return nil
}
