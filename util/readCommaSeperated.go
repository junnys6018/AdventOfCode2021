package util

import (
	"os"
	"strings"
)

func ReadCommaSeperated(file string) []int {
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return ListToInt(strings.Split(string(content), ","))
}
