package action

import (
	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/jira-work-log-sender/internal/jira"
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
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
