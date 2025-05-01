package import_data

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	mainInformationAnchor = "## "
	descriptionAnchor     = "- "
	tagAnchor             = "["
)

func splitFileToSections(file *os.File) ([][]string, error) {
	scanner := bufio.NewScanner(file)

	var sections [][]string
	var section []string
	lineNum := 2

	// Skip first two lines
	scanner.Scan()
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		lineNum++

		if line == "" {
			if len(section) > 0 {
				sections = append(sections, section)
				section = []string{}
			}
			continue
		}

		if isMainInformationLine(line) || isDescriptionLine(line) {
			section = append(section, line)
			continue
		}

		return nil, fmt.Errorf("impossible to parse line %d: %s", lineNum, line)
	}

	if len(section) > 0 {
		sections = append(sections, section)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error while reading file: %s", err)
	}

	return sections, nil
}

func isMainInformationLine(line string) bool {
	return strings.Index(line, mainInformationAnchor) == 0
}

func isDescriptionLine(line string) bool {
	trimmedLine := strings.TrimSpace(line)
	return strings.Index(trimmedLine, descriptionAnchor) == 0 || strings.Index(line, tagAnchor) == 0
}
