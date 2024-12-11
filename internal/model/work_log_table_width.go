package model

import (
	"strconv"
)

const (
	SendStatusWidth = 15
	TotalTimeWidth  = 15
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
	w := &WorkLogTableWidth{
		Number:       len(strconv.Itoa(len(workLogs))) + 1,
		HeaderText:   4,
		OriginalTime: 1,
		ModifiedTime: 2,
		IssueNumber:  5,
		Description:  11,
		SendStatus:   SendStatusWidth,
		TotalTime:    TotalTimeWidth,
	}

	for _, workLog := range workLogs {
		if len(workLog.GetHeader()) > w.HeaderText {
			w.HeaderText = len(workLog.GetHeader())
		}
		if len(workLog.OriginalTime.String()) > w.OriginalTime {
			w.OriginalTime = len(workLog.OriginalTime.String())
		}
		if len(workLog.ModifiedTime.String()) > w.ModifiedTime {
			w.ModifiedTime = len(workLog.ModifiedTime.String())
		}
		if len(workLog.IssueNumber) > w.IssueNumber {
			w.IssueNumber = len(workLog.IssueNumber)
		}
		if len(workLog.Description) > w.Description {
			w.Description = len(workLog.Description)
		}
	}

	width -= 20
	total := w.Number + w.HeaderText + w.OriginalTime + w.ModifiedTime + w.IssueNumber + w.Description
	if total > width {
		w.Description = w.Description - (total - width)
	}

	return w
}
