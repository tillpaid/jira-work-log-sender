package pages

import (
	"fmt"
	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/resource"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui/pages/page_dump_work_logs"
)

func DrawDumpWorkLogs(screen *goncurses.Window, config *resource.Config) error {
	if err := screen.Clear(); err != nil {
		return fmt.Errorf("error clearing screen: %v", err)
	}

	_, width := screen.MaxYX()
	rows := buildDumpWorkLogsRows(config, width)

	for i, line := range rows {
		screen.MovePrint(i, 0, prepareRow(line, width))
	}

	screen.Refresh()
	return nil
}

func buildDumpWorkLogsRows(config *resource.Config, width int) []string {
	delimiter := getDelimiter(width)
	var rows []string

	rows = append(rows, page_dump_work_logs.GetHeader(delimiter)...)
	rows = append(rows, page_dump_work_logs.GetBody(config)...)
	rows = append(rows, page_dump_work_logs.GetFooter(delimiter)...)

	return rows
}
