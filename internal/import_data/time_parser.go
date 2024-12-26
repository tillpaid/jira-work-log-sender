package import_data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/tillpaid/jira-work-log-sender/internal/model"
)

func parseTimeString(timeString string) (model.WorkLogTime, error) {
	var workLogTime model.WorkLogTime
	var err error

	if strings.Count(timeString, hoursChar) > 1 || strings.Count(timeString, minutesChar) > 1 {
		return workLogTime, errors.New("too many hours or minutes in time string")
	}

	workLogTime.Hours, err = parseTimeForChar(timeString, hoursChar)
	if err != nil {
		return workLogTime, err
	}

	if workLogTime.Hours > 0 {
		timeString = timeString[strings.Index(timeString, hoursChar)+1:]
	}

	workLogTime.Minutes, err = parseTimeForChar(timeString, minutesChar)
	if err != nil {
		return workLogTime, err
	}

	workLogTime.Hours += workLogTime.Minutes / 60
	workLogTime.Minutes = workLogTime.Minutes % 60

	if workLogTime.Hours == 0 && workLogTime.Minutes == 0 {
		return workLogTime, errors.New("no hours or minutes in time string")
	}

	return workLogTime, nil
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
