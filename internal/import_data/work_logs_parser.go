package import_data

import (
	"fmt"
	"os"

	"github.com/tillpaid/paysera-log-time-golang/internal/model"
	"github.com/tillpaid/paysera-log-time-golang/internal/resource"
	"github.com/tillpaid/paysera-log-time-golang/internal/service"
)

func ParseWorkLogs(config *resource.Config, oldWorkLogs []model.WorkLog) ([]model.WorkLog, error) {
	file, err := os.Open(config.PathToInputFile)
	if err != nil {
		return nil, fmt.Errorf("error opening file during parsing work logs: %v", err)
	}
	defer file.Close()

	var workLogs []model.WorkLog

	sections, err := splitFileToSections(file)
	if err != nil {
		return nil, fmt.Errorf("error splitting file to sections: %v", err)
	}

	for i, section := range sections {
		workLog, err := buildWorkLogFromSection(config, section, i+1)
		if err != nil {
			return nil, fmt.Errorf("error building work log: %v", err)
		}

		workLogs = append(workLogs, workLog)
	}

	workLogs = copyTempValuesFromOldWorkLogs(oldWorkLogs, workLogs)
	workLogs = service.ModifyWorkLogsTime(workLogs)

	return workLogs, nil
}

func copyTempValuesFromOldWorkLogs(oldWorkLogs []model.WorkLog, workLogs []model.WorkLog) []model.WorkLog {
	for _, oldWorkLog := range oldWorkLogs {
		if !oldWorkLog.ModifyTimeDisabled {
			continue
		}

		for i, workLog := range workLogs {
			if oldWorkLog.IssueNumber == workLog.IssueNumber && oldWorkLog.Description == workLog.Description {
				workLogs[i].ModifyTimeDisabled = true
			}
		}
	}

	return workLogs
}
