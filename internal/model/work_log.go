package model

import (
	"fmt"
	"strconv"
	"strings"
)

type WorkLog struct {
	Number       int
	OriginalTime string
	ModifiedTime string
	IssueNumber  string
	Description  string
}

func (w *WorkLog) ToStringWithSpaces(width *WorkLogTableWidth) string {
	return fmt.Sprintf(
		"%s %s | %s | %s | %s",
		getTextWithSpaces(strconv.Itoa(w.Number)+".", width.Number),
		getTextWithSpaces(w.OriginalTime, width.OriginalTime),
		getTextWithSpaces(w.ModifiedTime, width.ModifiedTime),
		getTextWithSpaces(w.IssueNumber, width.IssueNumber),
		w.Description,
	)
}

func getTextWithSpaces(text string, width int) string {
	neededSpaces := width - len(text)

	spaces := ""
	if neededSpaces > 0 {
		spaces = strings.Repeat(" ", neededSpaces)
	}

	return text + spaces
}
