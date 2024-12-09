package action

import (
	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/jira"
)

type Actions struct {
	PrintWorkLogs *PrintWorkLogsAction
	SendWorkLogs  *SendWorkLogsAction
}

func NewActions(client *jira.Client, window *goncurses.Window) *Actions {
	return &Actions{
		PrintWorkLogs: NewPrintWorkLogsAction(client, window),
		SendWorkLogs:  NewSendWorkLogsAction(client, window),
	}
}
