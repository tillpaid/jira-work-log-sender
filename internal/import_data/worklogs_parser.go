package import_data

import (
	"fmt"
	"os"

	"github.com/tillpaid/jira-work-log-sender/internal/model"
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
	"github.com/tillpaid/jira-work-log-sender/internal/service"
)

func ParseWorklogs(cfg *resource.Config, oldWorklogs []model.Worklog) ([]model.Worklog, error) {
	file, err := os.Open(cfg.Input.WorklogFile)
	if err != nil {
		return nil, fmt.Errorf("error opening file during parsing work logs: %v", err)
	}
	defer file.Close()

	var worklogs []model.Worklog

	sections, err := splitFileToSections(file)
	if err != nil {
		return nil, fmt.Errorf("error splitting file to sections: %v", err)
	}

	for i, section := range sections {
		worklog, err := buildWorklogFromSection(cfg, section, i+1)
		if err != nil {
			return nil, fmt.Errorf("error building work log: %v", err)
		}

		worklogs = append(worklogs, worklog)
	}

	worklogs = copyTempValuesFromOldWorklogs(oldWorklogs, worklogs)
	worklogs = service.ModifyWorklogsTime(worklogs, cfg)

	return worklogs, nil
}

func copyTempValuesFromOldWorklogs(oldWorklogs []model.Worklog, worklogs []model.Worklog) []model.Worklog {
	for _, oldWorklog := range oldWorklogs {
		if !oldWorklog.ModifyTimeDisabled {
			continue
		}

		for i, worklog := range worklogs {
			if oldWorklog.IssueNumber == worklog.IssueNumber && oldWorklog.Description == worklog.Description {
				worklogs[i].ModifyTimeDisabled = true
			}
		}
	}

	return worklogs
}
