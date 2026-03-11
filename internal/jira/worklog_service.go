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

type worklogService struct {
	currentUsername string
	jiraClient      *jira.Client
	cfg             *resource.Config
}

func newWorklogService(currentUsername string, jiraClient *jira.Client, cfg *resource.Config) *worklogService {
	return &worklogService{currentUsername: currentUsername, jiraClient: jiraClient, cfg: cfg}
}

func (w *worklogService) GetSpentTime(issueKey string) (*model.WorklogTime, error) {
	worklogs, _, err := w.jiraClient.Issue.GetWorklogs(issueKey)
	if err != nil {
		return nil, err
	}

	totalTimeSpent := &model.WorklogTime{}

	for _, worklog := range worklogs.Worklogs {
		if worklog.Author.EmailAddress != w.currentUsername {
			continue
		}

		totalTimeSpent.AddSeconds(worklog.TimeSpentSeconds)
	}

	return totalTimeSpent, nil
}

func (w *worklogService) SendWorklog(worklog model.Worklog) error {
	if w.cfg.IsDevRun {
		service.SleepMilliseconds(service.GetRandomInt(100, 500))
		return nil
	}

	if w.cfg.Tempo.UseTempoApiToSendWorklogs {
		return w.sendWorklogViaTempoApi(worklog)
	}

	return w.sendWorklogViaJiraApi(worklog)
}

func (w *worklogService) sendWorklogViaJiraApi(worklog model.Worklog) error {
	record := &jira.WorklogRecord{
		TimeSpent: worklog.ModifiedTime.String(),
		Comment:   fmt.Sprintf("%s\n%s", worklog.Tag, worklog.Description),
	}

	_, _, err := w.jiraClient.Issue.AddWorklogRecord(worklog.IssueNumber, record)

	return err
}

func (w *worklogService) sendWorklogViaTempoApi(worklog model.Worklog) error {
	data := TempoCreateWorklogRequest{
		Attributes: tempoCreateWorklogAttributes{
			EngineeringActivities: tempoCreateWorklogEngineeringActivitiesAttribute{
				Name:            w.cfg.Tempo.EngineeringActivityName,
				WorkAttributeId: w.cfg.Tempo.EngineeringActivityWorkAttributeID,
				Value:           strings.ReplaceAll(strings.Trim(worklog.Tag, "[]"), " ", ""),
			},
		},
		BillableSeconds:       nil,
		OriginId:              -1,
		Worker:                w.cfg.Tempo.WorkerID,
		Comment:               worklog.Description,
		Started:               time.Now().Format(time.DateOnly),
		TimeSpentSeconds:      worklog.ModifiedTime.GetInSeconds(),
		OriginTaskId:          worklog.IssueID,
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
