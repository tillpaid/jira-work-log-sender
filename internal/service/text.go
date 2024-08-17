package service

import "strings"

func GetTextWithSpaces(text string, width int) string {
	neededSpaces := width - len(text)

	spaces := ""
	if neededSpaces > 0 {
		spaces = strings.Repeat(" ", neededSpaces)
	}

	return text + spaces
}
