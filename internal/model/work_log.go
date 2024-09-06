package model

import (
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
	ModifiedTime WorkLogTime
	IssueNumber  string
	Description  string
}

func (wt *WorkLogTime) AddMinutes(minutes int) {
	wt.Hours += minutes / 60
	wt.Minutes += minutes % 60

	if wt.Minutes >= 60 {
		wt.Hours += wt.Minutes / 60
		wt.Minutes = wt.Minutes % 60
	}
}

func (wt *WorkLogTime) String() string {
	var parts []string

	if wt.Hours > 0 {
		parts = append(parts, strconv.Itoa(wt.Hours)+"h")
	}

	if wt.Minutes > 0 {
		parts = append(parts, strconv.Itoa(wt.Minutes)+"m")
	}

	return strings.Join(parts, " ")
}
