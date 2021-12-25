package util

import (
	"os"
	"strings"
)

func ReadTokens(file string) []string {
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return strings.Fields(string(content))
}
