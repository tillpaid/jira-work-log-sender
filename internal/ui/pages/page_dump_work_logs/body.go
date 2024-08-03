package page_dump_work_logs

import (
	"bufio"
	"fmt"
	"github.com/tillpaid/paysera-log-time-golang/internal/resource"
	"os"
	"strings"
)

func GetBody(config *resource.Config, delimiter string) ([]string, error) {
	outputFileContent, err := getOutputFileContent(config.OutputShellFile)
	if err != nil {
		return nil, err
	}

	relativeFilePath := strings.ReplaceAll(config.OutputShellFile, os.Getenv("HOME"), "~")

	output := []string{
		fmt.Sprintf("Successfully dumped work logs to file: %s", relativeFilePath),
		delimiter,
	}

	return append(output, outputFileContent...), nil
}

func getOutputFileContent(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
