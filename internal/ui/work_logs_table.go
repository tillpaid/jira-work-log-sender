package ui

import (
	"fmt"
	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/model"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui/work_logs_table"
	"strings"
)

func DrawWorkLogsTable(screen *goncurses.Window, workLogs []model.WorkLog) error {
	if err := screen.Clear(); err != nil {
		return fmt.Errorf("error clearing screen: %v", err)
	}

	_, width := screen.MaxYX()
	tableRows := buildTableRows(workLogs, width)

	for i, line := range tableRows {
		screen.MovePrint(i, 0, prepareRow(line, width))
	}

	screen.Refresh()
	return nil
}

func buildTableRows(workLogs []model.WorkLog, width int) []string {
	delimiter := fmt.Sprintf("+%s+", strings.Repeat("-", width-2))
	var rows []string

	rows = append(rows, work_logs_table.GetHeader(delimiter)...)
	rows = append(rows, work_logs_table.GetBody(workLogs)...)
	rows = append(rows, work_logs_table.GetFooter(workLogs, delimiter)...)

	return rows
}

func prepareRow(text string, width int) string {
	if text[0] == '+' {
		return text
	}

	textLen := len(text)
	spacesLen := width - 4 - textLen

	spaces := ""
	if spacesLen > 0 {
		spaces = strings.Repeat(" ", spacesLen)
	}

	if textLen > width-4 {
		text = text[:width-7] + "..."
	}

	return fmt.Sprintf("| %s%s |", text, spaces)
}
