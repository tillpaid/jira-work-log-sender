package export_data

import (
	"fmt"
	"os"

	"github.com/tillpaid/paysera-log-time-golang/internal/import_data"
	"github.com/tillpaid/paysera-log-time-golang/internal/jira"
	"github.com/tillpaid/paysera-log-time-golang/internal/resource"
)

func DumpWorkLogs(client *jira.Client, config *resource.Config) error {
	workLogs, err := import_data.ParseWorkLogs(client, config)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(config.OutputShellFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return fmt.Errorf("error opening file: %s", err)
	}
	defer file.Close()

	fileContent := convertWorkLogsToFileContent(workLogs)
	if _, err = file.WriteString(fileContent); err != nil {
		return fmt.Errorf("error writing to file: %s", err)
	}

	return nil
}
