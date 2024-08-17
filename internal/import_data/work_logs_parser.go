package import_data

import (
	"fmt"
	"os"

	"github.com/tillpaid/paysera-log-time-golang/internal/jira"
	"github.com/tillpaid/paysera-log-time-golang/internal/model"
	"github.com/tillpaid/paysera-log-time-golang/internal/resource"
	"github.com/tillpaid/paysera-log-time-golang/internal/service"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui/pages"
)

func ParseWorkLogs(loading *pages.Loading, client *jira.Client, config *resource.Config) ([]model.WorkLog, error) {
	loading.PrintRow("", 0)
	loading.PrintRow("Processing workLogs", 0)

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
		workLog, err := buildWorkLogFromSection(loading, client, config, section, i+1)
		if err != nil {
			return nil, fmt.Errorf("error building work log: %v", err)
		}

		workLogs = append(workLogs, workLog)
	}

	workLogs = service.ModifyWorkLogsTime(workLogs)

	return workLogs, nil
}
