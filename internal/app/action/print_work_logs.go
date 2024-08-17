package action

import (
	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/model"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui/pages"
)

func PrintWorkLogs(screen *goncurses.Window, workLogs []model.WorkLog) error {
	if err := pages.DrawWorkLogsTable(screen, workLogs); err != nil {
		return err
	}

	return nil
}
