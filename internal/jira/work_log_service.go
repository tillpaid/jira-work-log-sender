package jira

import (
	"fmt"
	"time"

	"github.com/andygrunwald/go-jira"
	"github.com/tillpaid/paysera-log-time-golang/internal/model"
)

type workLogService struct {
	currentUsername string
	jiraClient      *jira.Client
}

func newWorkLogService(currentUsername string, jiraClient *jira.Client) *workLogService {
	return &workLogService{currentUsername: currentUsername, jiraClient: jiraClient}
}

func (w *workLogService) GetSpentTime(issueID string) string {
	defaultAnswer := "n/a"

	workLogs, _, err := w.jiraClient.Issue.GetWorklogs(issueID)
	if err != nil {
		return defaultAnswer
	}

	var totalTimeSpent time.Duration

	for _, workLog := range workLogs.Worklogs {
		if workLog.Author.EmailAddress != w.currentUsername {
			continue
		}

		totalTimeSpent += time.Duration(workLog.TimeSpentSeconds) * time.Second
	}

	hours := int(totalTimeSpent.Hours())
	minutes := int(totalTimeSpent.Minutes()) % 60

	return fmt.Sprintf("%02dh %02dm", hours, minutes)
}

func (w *workLogService) SendWorkLog(workLog model.WorkLog) error {
	// todo: temp code
	time.Sleep(200 * time.Millisecond)
	return nil
	// todo: temp code

	record := &jira.WorklogRecord{
		TimeSpent: workLog.ModifiedTime.String(),
		Comment:   workLog.Description,
	}

	_, _, err := w.jiraClient.Issue.AddWorklogRecord(workLog.IssueNumber, record)

	return err
}
