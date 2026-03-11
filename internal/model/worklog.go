package model

import (
	"strconv"
	"strings"
)

type WorklogTime struct {
	Hours   int
	Minutes int
}

type Worklog struct {
	HeaderText                     string
	Number                         int
	OriginalTime                   WorklogTime
	ModifiedTime                   WorklogTime
	IssueNumber                    string
	IssueID                        string
	Tag                            string
	Description                    string
	ExcludedFromSpentTimeHighlight bool
	ModifyTimeDisabled             bool
}

func (w *Worklog) GetHeader() string {
	if len(w.HeaderText) > 30 {
		return w.HeaderText[:30]
	}

	return w.HeaderText
}

func (w *Worklog) ToggleModifyTime() {
	w.ModifyTimeDisabled = !w.ModifyTimeDisabled
}

func (wt *WorklogTime) AddMinutes(minutes int) {
	wt.Hours += minutes / 60
	wt.Minutes += minutes % 60

	if wt.Minutes >= 60 {
		wt.Hours += wt.Minutes / 60
		wt.Minutes = wt.Minutes % 60
	}
}

func (wt *WorklogTime) AddSeconds(seconds int) {
	wt.AddMinutes(seconds / 60)
}

func (wt *WorklogTime) GetInMinutes() int {
	return wt.Hours*60 + wt.Minutes
}

func (wt *WorklogTime) GetInSeconds() int {
	return wt.Hours*3600 + wt.Minutes*60
}

func (wt *WorklogTime) String() string {
	var parts []string

	if wt.Hours > 0 {
		parts = append(parts, strconv.Itoa(wt.Hours)+"h")
	}

	if wt.Minutes > 0 {
		parts = append(parts, strconv.Itoa(wt.Minutes)+"m")
	}

	if wt.Minutes == 0 && wt.Hours == 0 {
		parts = append(parts, "0m")
	}

	return strings.Join(parts, " ")
}
