package jira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/andygrunwald/go-jira"
	"github.com/tillpaid/jira-work-log-sender/internal/model"
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
	"github.com/tillpaid/jira-work-log-sender/internal/service"
)

type workLogService struct {
	currentUsername string
	jiraClient      *jira.Client
	cfg             *resource.Config
}

func newWorkLogService(currentUsername string, jiraClient *jira.Client, cfg *resource.Config) *workLogService {
	return &workLogService{currentUsername: currentUsername, jiraClient: jiraClient, cfg: cfg}
}

func (w *workLogService) GetSpentTime(issueKey string) (*model.WorkLogTime, error) {
	workLogs, _, err := w.jiraClient.Issue.GetWorklogs(issueKey)
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
	if w.cfg.IsDevRun {
		service.SleepMilliseconds(service.GetRandomInt(100, 500))
		return nil
	}

	if w.cfg.Tempo.UseTempoApiToSendWorklogs {
		return w.sendWorkLogViaTempoApi(workLog)
	}

	return w.sendWorkLogViaJiraApi(workLog)
}

func (w *workLogService) sendWorkLogViaJiraApi(workLog model.WorkLog) error {
	record := &jira.WorklogRecord{
		TimeSpent: workLog.ModifiedTime.String(),
		Comment:   fmt.Sprintf("%s\n%s", workLog.Tag, workLog.Description),
	}

	_, _, err := w.jiraClient.Issue.AddWorklogRecord(workLog.IssueNumber, record)

	return err
}

func (w *workLogService) sendWorkLogViaTempoApi(workLog model.WorkLog) error {
	data := TempoCreateWorklogRequest{
		Attributes: tempoCreateWorklogAttributes{
			EngineeringActivities: tempoCreateWorklogEngineeringActivitiesAttribute{
				Name:            w.cfg.Tempo.EngineeringActivityName,
				WorkAttributeId: w.cfg.Tempo.EngineeringActivityWorkAttributeID,
				Value:           strings.ReplaceAll(strings.Trim(workLog.Tag, "[]"), " ", ""),
			},
		},
		BillableSeconds:       nil,
		OriginId:              -1,
		Worker:                w.cfg.Tempo.WorkerID,
		Comment:               workLog.Description,
		Started:               time.Now().Format(time.DateOnly),
		TimeSpentSeconds:      workLog.ModifiedTime.GetInSeconds(),
		OriginTaskId:          workLog.IssueID,
		EndDate:               nil,
		IncludeNonWorkingDays: false,
	}

	url := w.cfg.Jira.Url + "/rest/tempo-timesheets/4/worklogs"
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(w.cfg.Jira.User, w.cfg.Jira.Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
