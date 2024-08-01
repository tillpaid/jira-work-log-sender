package ui

import (
	"fmt"
	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/model"
	"log"
	"strings"
)

func DrawTable(screen *goncurses.Window, width int, workLogs []model.WorkLog) {
	if err := screen.Clear(); err != nil {
		log.Fatalf("Error clearing screen: %v", err)
	}

	tableRows := buildTableRows(workLogs, width)

	for i, line := range tableRows {
		screen.MovePrint(i, 0, prepareRow(line, width))
	}

	screen.Refresh()
}

func buildTableRows(workLogs []model.WorkLog, width int) []string {
	delimiter := fmt.Sprintf("+%s+", strings.Repeat("-", width-2))
	menuTitle := "Log works for today"
	totalRow := "Total time: 6h 20m | Left: 1h 40m"

	workLogsTableWidth := model.NewWorkLogTableWidthWithCalculations(workLogs)
	rows := []string{delimiter, menuTitle, delimiter}

	for _, workLog := range workLogs {
		rows = append(rows, workLog.ToStringWithSpaces(workLogsTableWidth))
	}

	return append(rows, delimiter, totalRow, delimiter)
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
