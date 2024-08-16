package action

import (
	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/import_data"
	"github.com/tillpaid/paysera-log-time-golang/internal/jira"
	"github.com/tillpaid/paysera-log-time-golang/internal/resource"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui/pages"
)

func PrintWorkLogs(client *jira.Client, config *resource.Config, screen *goncurses.Window) error {
	workLogs, err := import_data.ParseWorkLogs(client, config)
	if err != nil {
		return err
	}

	if err = pages.DrawWorkLogsTable(screen, workLogs); err != nil {
		return err
	}

	return nil
}
