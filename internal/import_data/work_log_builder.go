package import_data

import (
	"errors"
	"fmt"
	"strings"

	"github.com/tillpaid/jira-work-log-sender/internal/model"
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
)

const (
	hoursChar   = "h"
	minutesChar = "m"
)

func buildWorkLogFromSection(config *resource.Config, section []string, number int) (model.WorkLog, error) {
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

	if !containAllowedTag(config, section[1]) {
		return workLog, fmt.Errorf("description does not contain allowed tags for task %s", issueNumber)
	}

	workLog.HeaderText = strings.TrimLeft(section[0], "# ")
	workLog.IssueNumber = issueNumber
	workLog.OriginalTime = originalTime
	workLog.Description = strings.TrimLeft(strings.Join(section[1:], "\n"), "- ")
	workLog.ExcludedFromSpentTimeHighlight = isExcludedFromTimeHighlight(workLog.IssueNumber, config)

	return workLog, nil
}

func getMainInformation(line string) (string, model.WorkLogTime, error) {
	mainParts := strings.Split(line, "|")
	if len(mainParts) != 2 {
		return "", model.WorkLogTime{}, errors.New("no pipe in main information or too many pipes")
	}

	parts := strings.Split(mainParts[1], " ")
	var secondParts []string

	for _, part := range parts {
		trimmedPart := strings.TrimSpace(part)
		if len(trimmedPart) > 0 {
			secondParts = append(secondParts, trimmedPart)
		}
	}

	switch len(secondParts) {
	case 0:
		return "", model.WorkLogTime{}, errors.New("no information after pipe")
	case 1:
		return secondParts[0], model.WorkLogTime{}, nil
	case 2:
		originalTime, err := parseTimeString(secondParts[1])
		if err != nil {
			return "", originalTime, fmt.Errorf("impossible to parse time string: %s", err)
		}

		return secondParts[0], originalTime, nil
	default:
		return "", model.WorkLogTime{}, errors.New("too many parts after pipe")
	}
}
