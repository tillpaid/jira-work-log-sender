package service

import "github.com/tillpaid/jira-work-log-sender/internal/ui"

func HandlePanic() {
	if r := recover(); r != nil {
		ui.EndWindow()
		panic(r)
	}
}
