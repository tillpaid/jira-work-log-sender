package import_data

import (
	"errors"
	"fmt"
	"github.com/tillpaid/paysera-log-time-golang/internal/model"
	"strings"
)

const (
	hoursChar   = "h"
	minutesChar = "m"
)

func buildWorkLogFromSection(section []string, number int) (model.WorkLog, error) {
	workLog := model.WorkLog{
		Number: number,
	}

	issueNumber, originalTime, err := getMainInformation(section[0])
	if err != nil {
		return workLog, fmt.Errorf("impossible to parse main information from section %d: %s", number, err)
	}

	workLog.IssueNumber = issueNumber
	workLog.OriginalTime = originalTime
	workLog.Description = strings.Join(section[1:], "\n")

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
