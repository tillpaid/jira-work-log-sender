package import_data

import (
	"errors"
	"fmt"
	"strings"

	"github.com/tillpaid/paysera-log-time-golang/internal/jira"
	"github.com/tillpaid/paysera-log-time-golang/internal/model"
	"github.com/tillpaid/paysera-log-time-golang/internal/resource"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui/pages"
)

const (
	hoursChar   = "h"
	minutesChar = "m"
)

func buildWorkLogFromSection(loading *pages.Loading, client *jira.Client, config *resource.Config, section []string, number int) (model.WorkLog, error) {
	workLog := model.WorkLog{
		Number: number,
	}

	issueNumber, originalTime, err := getMainInformation(section[0])
	if err != nil {
		return workLog, fmt.Errorf("impossible to parse main information from section %d: %s", number, err)
	}

	if len(section) < 2 {
		return workLog, fmt.Errorf("no description for task %s", issueNumber)
	}

	validDescription := checkWorkLogAllowedTag(config, section[1])
	if !validDescription {
		return workLog, fmt.Errorf("description does not contain allowed tags for task %s", issueNumber)
	}

	loading.PrintRow(fmt.Sprintf("Checking workLog %s in jira...", issueNumber), 0)

	issueExists, err := client.IssueService.IsIssueExists(issueNumber)
	if err != nil {
		return workLog, fmt.Errorf("impossible to check issue %s in jira: %s", issueNumber, err)
	}
	if !issueExists {
		return workLog, fmt.Errorf("issue %s does not exist", issueNumber)
	}

	workLog.IssueNumber = issueNumber
	workLog.OriginalTime = originalTime
	workLog.Description = strings.TrimLeft(strings.Join(section[1:], "\n"), "- ")

	return workLog, nil
}

func getMainInformation(line string) (string, model.WorkLogTime, error) {
	dashIndex := strings.Index(line, "-")
	pipeIndex := strings.Index(line, "|")
	lastSpaceRelativeIndex := strings.Index(line[pipeIndex+2:], " ")
	lastSpaceIndex := pipeIndex + 2 + lastSpaceRelativeIndex

	if dashIndex == -1 || pipeIndex == -1 || lastSpaceRelativeIndex == -1 {
		return "", model.WorkLogTime{}, errors.New("no dash or pipe or space after pipe")
	}

	issueNumber := strings.TrimSpace(line[pipeIndex+2 : lastSpaceIndex])

	originalTime, err := parseTimeString(strings.TrimSpace(line[lastSpaceIndex+1:]))
	if err != nil {
		return "", model.WorkLogTime{}, fmt.Errorf("impossible to parse time string: %s", err)
	}

	return issueNumber, originalTime, nil
}
