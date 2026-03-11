package action

import (
	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/jira-work-log-sender/internal/jira"
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
)

type Actions struct {
	PrintWorklogs *PrintWorklogsAction
	SendWorklogs  *SendWorklogsAction
}

func NewActions(client *jira.Client, window *goncurses.Window, cfg *resource.Config) *Actions {
	return &Actions{
		PrintWorklogs: NewPrintWorklogsAction(client, window, cfg),
		SendWorklogs:  NewSendWorklogsAction(client, window, cfg),
	}
}
