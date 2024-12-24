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

type WorkLogTableWidth struct {
	Number       int
	HeaderText   int
	OriginalTime int
	ModifiedTime int
	IssueNumber  int
	Description  int
	SendStatus   int
	TotalTime    int
}

func NewWorkLogTableWidthWithCalculations(workLogs []WorkLog, width int) *WorkLogTableWidth {
	width -= 6 // box borders and table dividers

	w := &WorkLogTableWidth{
		Number:       len(strconv.Itoa(len(workLogs))) + 3,
		HeaderText:   headerTextFirstMinWidth,
		OriginalTime: 4,
		ModifiedTime: 4,
		IssueNumber:  11,
		Description:  descriptionFirstMinWidth,
		SendStatus:   sendStatusFirstMinWidth,
		TotalTime:    totalTimeFirstMinWidth,
	}

	for _, workLog := range workLogs {
		w.resolveWidthForWorkLog(workLog)
	}

	w.adjustWidthForWorkLogTable(width)
	w.adjustWidthForSendTable(width)

	return w
}

func (w *WorkLogTableWidth) getTotalForWorkLogTable() int {
	return w.Number + w.HeaderText + w.OriginalTime + w.ModifiedTime + w.IssueNumber + w.Description
}

func (w *WorkLogTableWidth) getTotalForSendTable() int {
	return w.Number + w.IssueNumber + w.SendStatus + w.TotalTime
}

func (w *WorkLogTableWidth) resolveWidthForWorkLog(workLog WorkLog) {
	w.resolveWidth(&w.HeaderText, workLog.GetHeader())
	w.resolveWidth(&w.OriginalTime, workLog.OriginalTime.String())
	w.resolveWidth(&w.ModifiedTime, workLog.ModifiedTime.String())
	w.resolveWidth(&w.IssueNumber, workLog.IssueNumber)
	w.resolveWidth(&w.Description, workLog.Description)
}

func (w *WorkLogTableWidth) resolveWidth(currentWidth *int, text string) {
	if len(text)+2 > *currentWidth {
		*currentWidth = len(text) + 2
	}
}

func (w *WorkLogTableWidth) adjustWidthForWorkLogTable(width int) {
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
		total := w.getTotalForWorkLogTable()
		if total <= width {
			break
		}

		reduceWidth(adj.field, adj.minWidth, total-width)
	}
}

func (w *WorkLogTableWidth) adjustWidthForSendTable(width int) {
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
