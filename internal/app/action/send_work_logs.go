package action

import (
	"fmt"
	"strings"

	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/jira-work-log-sender/internal/jira"
	"github.com/tillpaid/jira-work-log-sender/internal/model"
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
	"github.com/tillpaid/jira-work-log-sender/internal/service"
	"github.com/tillpaid/jira-work-log-sender/internal/ui"
	"github.com/tillpaid/jira-work-log-sender/internal/ui/element/table"
	"github.com/tillpaid/jira-work-log-sender/internal/ui/pages/page_send_work_logs"
)

const (
	toWaiting = iota
	toSending
	toDone
	toFailed
	toCalculating
	toCustomText
)

const (
	statusIndex = 3
	timeIndex   = 4
)

const (
	textWaiting     = "Waiting..."
	textSending     = "Sending..."
	textDone        = "Done!"
	textFailed      = "Failed!"
	textCalculating = "Calculating..."
	emptyText       = ""
)

type Transition struct {
	Previous string
	Next     string
	Color    int16
}

func (st *Transition) GetText() string {
	if len(st.Previous) > len(st.Next) {
		return st.Next + strings.Repeat(" ", len(st.Previous)-len(st.Next))
	}

	return st.Next
}

var transitions = map[int16]*Transition{
	toWaiting:     {emptyText, textWaiting, ui.CyanOnBlack},
	toSending:     {textWaiting, textSending, ui.MagentaOnBlack},
	toDone:        {textSending, textDone, ui.GreenOnBlack},
	toFailed:      {textSending, textFailed, ui.RedOnBlack},
	toCalculating: {textWaiting, textCalculating, ui.MagentaOnBlack},
	toCustomText:  {textCalculating, emptyText, ui.DefaultColor},
}

type SendWorkLogsAction struct {
	client *jira.Client
	window *goncurses.Window
	config *resource.Config
}

func NewSendWorkLogsAction(client *jira.Client, window *goncurses.Window, config *resource.Config) *SendWorkLogsAction {
	return &SendWorkLogsAction{client: client, window: window, config: config}
}

func (a *SendWorkLogsAction) Send(workLogs []model.WorkLog) error {
	if err := a.client.Ping(); err != nil {
		return err
	}

	for _, workLog := range workLogs {
		if workLog.OriginalTime.Hours == 0 && workLog.OriginalTime.Minutes == 0 {
			return fmt.Errorf("work log with issue number %s has no time spent", workLog.IssueNumber)
		}
	}

	t, err := page_send_work_logs.DrawSendWorkLogsPage(a.window, workLogs)
	if err != nil {
		return err
	}

	for i := range workLogs {
		service.SleepMilliseconds(50)
		a.applyTransition(t, i, timeIndex, transitions[toWaiting])
	}

	for i := range workLogs {
		service.SleepMilliseconds(50)
		a.applyTransition(t, len(workLogs)-1-i, statusIndex, transitions[toWaiting])
	}

	for i, workLog := range workLogs {
		a.sendWorkLog(t, workLog, i)
	}

	for i, workLog := range workLogs {
		a.setSpentTime(t, workLog, i)
	}

	return nil
}

func (a *SendWorkLogsAction) sendWorkLog(table *table.Table, workLog model.WorkLog, i int) {
	a.applyTransition(table, i, statusIndex, transitions[toSending])

	if err := a.client.WorkLogService.SendWorkLog(workLog); err != nil {
		a.applyTransition(table, i, statusIndex, transitions[toFailed])
		return
	}

	a.applyTransition(table, i, statusIndex, transitions[toDone])
}

func (a *SendWorkLogsAction) setSpentTime(table *table.Table, workLog model.WorkLog, i int) {
	a.applyTransition(table, i, timeIndex, transitions[toCalculating])

	transitions[toCustomText].Next = "n/a"
	transitions[toCustomText].Color = ui.DefaultColor

	workLogTime, err := a.client.WorkLogService.GetSpentTime(workLog.IssueNumber)
	if err != nil {
		a.applyTransition(table, i, timeIndex, transitions[toCustomText])
		return
	}

	transitions[toCustomText].Next = workLogTime.String()
	if service.ShouldHighlightTimeForWorkLog(workLog, workLogTime, a.config) {
		transitions[toCustomText].Color = ui.YellowOnBlack
	}

	a.applyTransition(table, i, timeIndex, transitions[toCustomText])
}

func (a *SendWorkLogsAction) applyTransition(table *table.Table, rowI int, columnI int, transition *Transition) {
	row := table.Rows[rowI]
	row.Columns[columnI].Text = transition.GetText()
	row.Columns[columnI].Color = transition.Color

	table.ReDrawRow(row)
	a.window.Refresh()
}
