package action

import (
	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/export_data"
	"github.com/tillpaid/paysera-log-time-golang/internal/resource"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui/pages"
)

func DumpWorkLogs(config *resource.Config, screen *goncurses.Window) error {
	if err := export_data.DumpWorkLogs(config); err != nil {
		return err
	}

	if err := pages.DrawDumpWorkLogs(screen, config); err != nil {
		return err
	}

	return nil
}
