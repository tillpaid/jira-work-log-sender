package import_data

import (
	"fmt"
	"github.com/tillpaid/paysera-log-time-golang/internal/model"
	"github.com/tillpaid/paysera-log-time-golang/internal/resource"
	"github.com/tillpaid/paysera-log-time-golang/internal/service"
	"os"
)

func ParseWorkLogs(config *resource.Config) ([]model.WorkLog, error) {
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
		workLog, err := buildWorkLogFromSection(section, i+1)
		if err != nil {
			return nil, fmt.Errorf("error building work log: %v", err)
		}

		workLogs = append(workLogs, workLog)
	}

	workLogs = service.ModifyWorkLogsTime(workLogs)

	return workLogs, nil
}
