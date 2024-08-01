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

func (w *WorkLog) ToStringWithSpaces(width *WorkLogTableWidth, screenWidth int) string {
	number := getTextWithSpaces(strconv.Itoa(w.Number)+".", width.Number)
	originalTime := getTextWithSpaces(w.OriginalTime, width.OriginalTime)
	modifiedTime := getTextWithSpaces(w.ModifiedTime, width.ModifiedTime)
	issueNumber := getTextWithSpaces(w.IssueNumber, width.IssueNumber)

	return fmt.Sprintf("%s %s | %s | %s | %s",
		number, originalTime, modifiedTime, issueNumber, w.Description)
}

func getTextWithSpaces(text string, width int) string {
	neededSpaces := width - len(text)

	spaces := ""
	if neededSpaces > 0 {
		spaces = strings.Repeat(" ", neededSpaces)
	}

	return text + spaces
}
