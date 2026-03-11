package import_data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/tillpaid/jira-work-log-sender/internal/model"
)

func parseTimeString(timeString string) (model.WorklogTime, error) {
	var worklogTime model.WorklogTime
	var err error

	if strings.Count(timeString, hoursChar) > 1 || strings.Count(timeString, minutesChar) > 1 {
		return worklogTime, errors.New("too many hours or minutes in time string")
	}

	worklogTime.Hours, err = parseTimeForChar(timeString, hoursChar)
	if err != nil {
		return worklogTime, err
	}

	if worklogTime.Hours > 0 {
		timeString = timeString[strings.Index(timeString, hoursChar)+1:]
	}

	worklogTime.Minutes, err = parseTimeForChar(timeString, minutesChar)
	if err != nil {
		return worklogTime, err
	}

	worklogTime.Hours += worklogTime.Minutes / 60
	worklogTime.Minutes = worklogTime.Minutes % 60

	if worklogTime.Hours == 0 && worklogTime.Minutes == 0 {
		return worklogTime, errors.New("no hours or minutes in time string")
	}

	return worklogTime, nil
}

func parseTimeForChar(timeString string, char string) (int, error) {
	index := strings.Index(timeString, char)
	if index == -1 {
		return 0, nil
	}

	parts := strings.Split(timeString[:index], char)
	if len(parts) > 2 {
		return 0, fmt.Errorf("invalid value for %s in time string", char)
	}

	value, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("cannot parse value for %s in time string: %s", char, err)
	}

	if value < 0 {
		return 0, fmt.Errorf("value for %s in time string is negative", char)
	}

	return value, nil
}
