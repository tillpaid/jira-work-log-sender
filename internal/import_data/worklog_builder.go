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

func buildWorklogFromSection(cfg *resource.Config, section []string, number int) ([]model.Worklog, error) {
	issueNumberRaw, originalTime, err := getMainInformation(section[0])
	if err != nil {
		return nil, fmt.Errorf("impossible to parse main information from section %d: %s", number, err)
	}

	issueNumbers := splitIssueNumbers(issueNumberRaw)

	for _, issueNumber := range issueNumbers {
		if isForbiddenProject(issueNumber, cfg) {
			return nil, fmt.Errorf("task %s belongs to a forbidden project", issueNumber)
		}
	}

	if len(section) < 2 {
		return nil, fmt.Errorf("no description for task %s", issueNumberRaw)
	}

	tag := section[1]
	if !containAllowedTag(cfg, tag) {
		return nil, fmt.Errorf("description does not contain allowed tags for task %s", issueNumberRaw)
	}

	headerText := strings.TrimLeft(section[0], "# ")
	description := strings.Join(section[2:], "\n")

	var worklogs []model.Worklog
	for _, issueNumber := range issueNumbers {
		worklogs = append(worklogs, model.Worklog{
			Number:                         number,
			HeaderText:                     headerText,
			IssueNumber:                    issueNumber,
			OriginalTime:                   originalTime,
			Description:                    description,
			Tag:                            tag,
			ExcludedFromSpentTimeHighlight: isExcludedFromTimeHighlight(issueNumber, cfg),
		})
	}

	return worklogs, nil
}

func splitIssueNumbers(raw string) []string {
	parts := strings.Split(raw, ",")
	var result []string
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

func getMainInformation(line string) (string, model.WorklogTime, error) {
	mainParts := strings.Split(line, "|")
	if len(mainParts) != 2 {
		return "", model.WorklogTime{}, errors.New("no pipe in main information or too many pipes")
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
		return "", model.WorklogTime{}, errors.New("no information after pipe")
	case 1:
		return secondParts[0], model.WorklogTime{}, nil
	case 2:
		originalTime, err := parseTimeString(secondParts[1])
		if err != nil {
			return "", originalTime, fmt.Errorf("impossible to parse time string: %s", err)
		}

		return secondParts[0], originalTime, nil
	default:
		return "", model.WorklogTime{}, errors.New("too many parts after pipe")
	}
}

func isForbiddenProject(issueNumber string, cfg *resource.Config) bool {
	for _, forbiddenProject := range cfg.ForbiddenProjects {
		if strings.HasPrefix(issueNumber, forbiddenProject) {
			return true
		}
	}

	return false
}
