package action

import (
	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/model"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui/pages/page_work_logs"
)

type PrintWorkLogsAction struct {
	screen *goncurses.Window
}

func NewPrintWorkLogsAction(screen *goncurses.Window) *PrintWorkLogsAction {
	return &PrintWorkLogsAction{screen: screen}
}

func (a *PrintWorkLogsAction) Print(workLogs []model.WorkLog, selectedRow int) error {
	return page_work_logs.DrawWorkLogsTable(a.screen, workLogs, selectedRow)
}
