package model

import (
	"fmt"
	"strconv"
	"strings"
)

type WorkLogTime struct {
	Hours   int
	Minutes int
}

type WorkLog struct {
	Number       int
	OriginalTime WorkLogTime
	ModifiedTime string
	IssueNumber  string
	Description  string
}

func (w *WorkLog) ToStringWithSpaces(width *WorkLogTableWidth) string {
	return fmt.Sprintf(
		"%s %s | %s | %s | %s",
		w.getTextWithSpaces(strconv.Itoa(w.Number)+".", width.Number),
		w.getTextWithSpaces(w.OriginalTime.String(), width.OriginalTime),
		w.getTextWithSpaces(w.ModifiedTime, width.ModifiedTime),
		w.getTextWithSpaces(w.IssueNumber, width.IssueNumber),
		strings.ReplaceAll(w.Description, "\n", " "),
	)
}

func (w *WorkLog) getTextWithSpaces(text string, width int) string {
	neededSpaces := width - len(text)

	spaces := ""
	if neededSpaces > 0 {
		spaces = strings.Repeat(" ", neededSpaces)
	}

	return text + spaces
}

func (wt *WorkLogTime) String() string {
	var result string

	if wt.Hours > 0 {
		result += strconv.Itoa(wt.Hours) + "h"
	}

	if wt.Minutes > 0 {
		result += strconv.Itoa(wt.Minutes) + "m"
	}

	return result
}
