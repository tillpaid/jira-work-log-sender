package action

import (
	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/jira"
)

type Actions struct {
	PrintWorkLogs *PrintWorkLogsAction
	SendWorkLogs  *SendWorkLogsAction
}

func NewActions(client *jira.Client, screen *goncurses.Window) *Actions {
	return &Actions{
		PrintWorkLogs: NewPrintWorkLogsAction(client, screen),
		SendWorkLogs:  NewSendWorkLogsAction(client, screen),
	}
}
