package action

import (
	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/jira"
	"github.com/tillpaid/paysera-log-time-golang/internal/resource"
)

type Actions struct {
	PrintWorkLogs *PrintWorkLogsAction
	SendWorkLogs  *SendWorkLogsAction
}

func NewActions(client *jira.Client, window *goncurses.Window, config *resource.Config) *Actions {
	return &Actions{
		PrintWorkLogs: NewPrintWorkLogsAction(client, window),
		SendWorkLogs:  NewSendWorkLogsAction(client, window, config),
	}
}
