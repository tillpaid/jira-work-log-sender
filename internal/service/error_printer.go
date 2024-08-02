package service

import (
	"fmt"
	"os"
	"strings"
)

func PrintFatalError(err error) {
	parts := strings.Split(err.Error(), ": ")

	for i, part := range parts {
		var indent string
		if i > 0 {
			indent = fmt.Sprintf("|%s ", strings.Repeat("--", i))
		}

		fmt.Printf("\033[31m%s%s\033[0m\n", indent, part)
	}

	os.Exit(1)
}
