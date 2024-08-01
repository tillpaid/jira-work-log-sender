package model

import "strconv"

type WorkLogTableWidth struct {
	Number       int
	OriginalTime int
	ModifiedTime int
	IssueNumber  int
}

func NewWorkLogTableWidthWithCalculations(workLogs []WorkLog) *WorkLogTableWidth {
	w := &WorkLogTableWidth{
		Number: len(strconv.Itoa(len(workLogs))) + 1,
	}

	for _, workLog := range workLogs {
		if len(workLog.OriginalTime) > w.OriginalTime {
			w.OriginalTime = len(workLog.OriginalTime)
		}
		if len(workLog.ModifiedTime) > w.ModifiedTime {
			w.ModifiedTime = len(workLog.ModifiedTime)
		}
		if len(workLog.IssueNumber) > w.IssueNumber {
			w.IssueNumber = len(workLog.IssueNumber)
		}
	}

	return w
}
