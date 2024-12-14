package jira

import (
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

func (w *workLogService) GetSpentTime(issueID string) (*model.WorkLogTime, error) {
	workLogs, _, err := w.jiraClient.Issue.GetWorklogs(issueID)
	if err != nil {
		return nil, err
	}

	totalTimeSpent := &model.WorkLogTime{}

	for _, workLog := range workLogs.Worklogs {
		if workLog.Author.EmailAddress != w.currentUsername {
			continue
		}

		totalTimeSpent.AddSeconds(workLog.TimeSpentSeconds)
	}

	return totalTimeSpent, nil
}

func (w *workLogService) SendWorkLog(workLog model.WorkLog) error {
	record := &jira.WorklogRecord{
		TimeSpent: workLog.ModifiedTime.String(),
		Comment:   workLog.Description,
	}

	_, _, err := w.jiraClient.Issue.AddWorklogRecord(workLog.IssueNumber, record)

	return err
}
