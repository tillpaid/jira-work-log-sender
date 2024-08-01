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

	delimiter := fmt.Sprintf("+%s+", strings.Repeat("-", width-2))
	workLogsTable := buildWorkLogsTable(workLogs)

	// make menu title blue
	menuTitle := getRow("Log works for today", width)
	totalRow := getRow("Total time: 6h 20m | Left: 1h 40m", width)

	rows := []string{
		delimiter,
		menuTitle,
		delimiter,
	}

	for _, row := range workLogsTable {
		rows = append(rows, getRow(row, width))
	}

	rows = append(rows, delimiter)
	rows = append(rows, totalRow)
	rows = append(rows, delimiter)

	for i, line := range rows {
		screen.MovePrint(i, 0, line)
	}
	screen.Refresh()
}

func buildWorkLogsTable(workLogs []model.WorkLog) []string {
	var rows []string
	workLogsTableWidth := model.NewWorkLogTableWidthWithCalculations(workLogs)

	for _, workLog := range workLogs {
		rows = append(rows, workLog.ToStringWithSpaces(workLogsTableWidth))
	}

	return rows
}

func getRow(text string, width int) string {
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
