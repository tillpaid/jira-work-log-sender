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
	"github.com/tillpaid/jira-work-log-sender/internal/ui/pages/page_send_worklogs"
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

type SendWorklogsAction struct {
	client *jira.Client
	window *goncurses.Window
	cfg    *resource.Config
}

func NewSendWorklogsAction(client *jira.Client, window *goncurses.Window, cfg *resource.Config) *SendWorklogsAction {
	return &SendWorklogsAction{client: client, window: window, cfg: cfg}
}

func (a *SendWorklogsAction) Send(worklogs []model.Worklog) error {
	if err := a.client.Ping(); err != nil {
		return err
	}

	for _, worklog := range worklogs {
		if worklog.OriginalTime.Hours == 0 && worklog.OriginalTime.Minutes == 0 {
			return fmt.Errorf("work log with issue number %s has no time spent", worklog.IssueNumber)
		}
	}

	t, err := page_send_worklogs.DrawSendWorklogsPage(a.window, worklogs)
	if err != nil {
		return err
	}

	for i := range worklogs {
		service.SleepMilliseconds(50)
		a.applyTransition(t, i, timeIndex, transitions[toWaiting])
	}

	for i := range worklogs {
		service.SleepMilliseconds(50)
		a.applyTransition(t, len(worklogs)-1-i, statusIndex, transitions[toWaiting])
	}

	for i, worklog := range worklogs {
		a.sendWorklog(t, worklog, i)
	}

	for i, worklog := range worklogs {
		a.setSpentTime(t, worklog, i)
	}

	return nil
}

func (a *SendWorklogsAction) sendWorklog(table *table.Table, worklog model.Worklog, i int) {
	a.applyTransition(table, i, statusIndex, transitions[toSending])

	if err := a.client.WorklogService.SendWorklog(worklog); err != nil {
		a.applyTransition(table, i, statusIndex, transitions[toFailed])
		return
	}

	a.applyTransition(table, i, statusIndex, transitions[toDone])
}

func (a *SendWorklogsAction) setSpentTime(table *table.Table, worklog model.Worklog, i int) {
	a.applyTransition(table, i, timeIndex, transitions[toCalculating])

	transitions[toCustomText].Next = "n/a"
	transitions[toCustomText].Color = ui.DefaultColor

	worklogTime, err := a.client.WorklogService.GetSpentTime(worklog.IssueNumber)
	if err != nil {
		a.applyTransition(table, i, timeIndex, transitions[toCustomText])
		return
	}

	transitions[toCustomText].Next = worklogTime.String()
	if service.ShouldHighlightTimeForWorklog(worklog, worklogTime, a.cfg) {
		transitions[toCustomText].Color = ui.YellowOnBlack
	}

	a.applyTransition(table, i, timeIndex, transitions[toCustomText])
}

func (a *SendWorklogsAction) applyTransition(table *table.Table, rowI int, columnI int, transition *Transition) {
	row := table.Rows[rowI]
	row.Columns[columnI].Text = transition.GetText()
	row.Columns[columnI].Color = transition.Color

	table.ReDrawRow(row)
	a.window.Refresh()
}
