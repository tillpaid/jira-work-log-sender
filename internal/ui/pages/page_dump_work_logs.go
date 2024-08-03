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

	height, width := screen.MaxYX()
	rows, err := buildDumpWorkLogsRows(config, height, width)
	if err != nil {
		return err
	}

	for i, line := range rows {
		screen.MovePrint(i, 0, prepareRow(line, width))
	}

	screen.Refresh()
	return nil
}

func buildDumpWorkLogsRows(config *resource.Config, height int, width int) ([]string, error) {
	delimiter := getDelimiter(width)
	var rows []string

	header := page_dump_work_logs.GetHeader(delimiter)
	footer := page_dump_work_logs.GetFooter(delimiter)

	body, err := page_dump_work_logs.GetBody(config, delimiter)
	if err != nil {
		return nil, err
	}

	otherRowsLen := len(header) + len(footer)
	body = cutBody(body, otherRowsLen, height, width)

	rows = append(rows, header...)
	rows = append(rows, body...)
	rows = append(rows, footer...)

	return rows, nil
}
