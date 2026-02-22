package jira

import (
	"fmt"
	"log"

	"github.com/andygrunwald/go-jira"
	"github.com/tillpaid/jira-work-log-sender/internal/model"
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
	"github.com/tillpaid/jira-work-log-sender/internal/service"
)

type workLogService struct {
	currentUsername string
	jiraClient      *jira.Client
	config          *resource.Config
}

func newWorkLogService(currentUsername string, jiraClient *jira.Client, config *resource.Config) *workLogService {
	return &workLogService{currentUsername: currentUsername, jiraClient: jiraClient, config: config}
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
	if w.config.IsDevRun {
		service.SleepMilliseconds(service.GetRandomInt(100, 500))
		return nil
	}

	if w.config.Tempo.UseTempoApiToSendWorklogs {
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
	//payload := TempoCreateWorklogRequest{
	//	Attributes: tempoCreateWorklogAttributes{
	//		EngineeringActivities: tempoCreateWorklogEngineeringActivitiesAttribute{
	//			Name:            w.config.Tempo.EngineeringActivityName,
	//			WorkAttributeId: w.config.Tempo.EngineeringActivityWorkAttributeID,
	//			Value:           strings.Trim(workLog.Tag, "[]"),
	//		},
	//	},
	//	BillableSeconds:       nil,
	//	OriginId:              -1,
	//	Worker:                w.config.Tempo.WorkerID,
	//	Comment:               workLog.Description,
	//	Started:               time.Now().Format(time.DateOnly),
	//	TimeSpentSeconds:      workLog.ModifiedTime.GetInSeconds(),
	//	OriginTaskId:          workLog.IssueID,
	//	RemainingEstimate:     "",
	//	EndDate:               nil,
	//	IncludeNonWorkingDays: false,
	//}

	log.Fatal("Tempo API is not implemented yet")
	return nil
}
