package action

import (
	"os"
	"os/exec"

	"github.com/tillpaid/paysera-log-time-golang/internal/export_data"
	"github.com/tillpaid/paysera-log-time-golang/internal/jira"
	"github.com/tillpaid/paysera-log-time-golang/internal/resource"
)

func DumpWorkLogs(client *jira.Client, config *resource.Config) error {
	if err := export_data.DumpWorkLogs(client, config); err != nil {
		return err
	}

	cmd := exec.Command(config.OutputShellFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		_, err = os.Stderr.WriteString(err.Error())
		if err != nil {
			return err
		}
	}

	return nil
}
