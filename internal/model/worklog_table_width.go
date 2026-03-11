package model

import (
	"strconv"
)

const (
	headerTextFirstMinWidth   = 6
	headerTextSecondMinWidth  = 4
	descriptionFirstMinWidth  = 13
	descriptionSecondMinWidth = 4
	sendStatusFirstMinWidth   = 15
	sendStatusSecondMinWidth  = 8
	totalTimeFirstMinWidth    = 15
	totalTimeSecondMinWidth   = 8
)

type WorklogTableWidth struct {
	Number       int
	HeaderText   int
	OriginalTime int
	ModifiedTime int
	IssueNumber  int
	Description  int
	SendStatus   int
	TotalTime    int
}

func NewWorklogTableWidthWithCalculations(worklogs []Worklog, width int) *WorklogTableWidth {
	width -= 6 // box borders and table dividers

	w := &WorklogTableWidth{
		Number:       len(strconv.Itoa(len(worklogs))) + 3,
		HeaderText:   headerTextFirstMinWidth,
		OriginalTime: 4,
		ModifiedTime: 4,
		IssueNumber:  11,
		Description:  descriptionFirstMinWidth,
		SendStatus:   sendStatusFirstMinWidth,
		TotalTime:    totalTimeFirstMinWidth,
	}

	for _, worklog := range worklogs {
		w.resolveWidthForWorklog(worklog)
	}

	w.adjustWidthForWorklogTable(width)
	w.adjustWidthForSendTable(width)

	return w
}

func (w *WorklogTableWidth) getTotalForWorklogTable() int {
	return w.Number + w.HeaderText + w.OriginalTime + w.ModifiedTime + w.IssueNumber + w.Description
}

func (w *WorklogTableWidth) getTotalForSendTable() int {
	return w.Number + w.IssueNumber + w.SendStatus + w.TotalTime
}

func (w *WorklogTableWidth) resolveWidthForWorklog(worklog Worklog) {
	w.resolveWidth(&w.HeaderText, worklog.GetHeader())
	w.resolveWidth(&w.OriginalTime, worklog.OriginalTime.String())
	w.resolveWidth(&w.ModifiedTime, worklog.ModifiedTime.String())
	w.resolveWidth(&w.IssueNumber, worklog.IssueNumber)
	w.resolveWidth(&w.Description, worklog.Description)
}

func (w *WorklogTableWidth) resolveWidth(currentWidth *int, text string) {
	if len(text)+2 > *currentWidth {
		*currentWidth = len(text) + 2
	}
}

func (w *WorklogTableWidth) adjustWidthForWorklogTable(width int) {
	type widthAdjustment struct {
		field    *int
		minWidth int
	}

	adjustments := []widthAdjustment{
		{&w.Description, descriptionFirstMinWidth},
		{&w.HeaderText, headerTextSecondMinWidth},
		{&w.Description, descriptionSecondMinWidth},
	}

	for _, adj := range adjustments {
		total := w.getTotalForWorklogTable()
		if total <= width {
			break
		}

		reduceWidth(adj.field, adj.minWidth, total-width)
	}
}

func (w *WorklogTableWidth) adjustWidthForSendTable(width int) {
	if w.getTotalForSendTable() > width {
		w.SendStatus = sendStatusSecondMinWidth
		w.TotalTime = totalTimeSecondMinWidth
	}
}

func reduceWidth(current *int, minWidth, difference int) {
	*current -= difference
	if *current < minWidth {
		*current = minWidth
	}
}
