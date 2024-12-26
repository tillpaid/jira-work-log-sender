package jira

import (
	"math/rand"
	"time"

	"github.com/andygrunwald/go-jira"
	"github.com/tillpaid/jira-work-log-sender/internal/model"
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
)

type workLogService struct {
	currentUsername string
	jiraClient      *jira.Client
	config          *resource.Config
}

func newWorkLogService(currentUsername string, jiraClient *jira.Client, config *resource.Config) *workLogService {
	return &workLogService{currentUsername: currentUsername, jiraClient: jiraClient, config: config}
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
	if w.config.IsDevRun {
		time.Sleep(time.Duration(rand.Intn(500-100)+100) * time.Millisecond)
		return nil
	}

	record := &jira.WorklogRecord{
		TimeSpent: workLog.ModifiedTime.String(),
		Comment:   workLog.Description,
	}

	_, _, err := w.jiraClient.Issue.AddWorklogRecord(workLog.IssueNumber, record)

	return err
}
