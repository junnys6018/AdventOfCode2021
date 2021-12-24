package day01

import (
	"fmt"
	"os"
	"strings"

	"advent/util"
)

func Answer() {
	content, err := os.ReadFile("day01/input")
	if err != nil {
		panic(err)
	}

	depthsStr := strings.Fields(string(content))
	depthsInt := make([]int, len(depthsStr))
	for i := range depthsStr {
		depthsInt[i] = util.ReadInt(depthsStr[i])
	}

	count := 0
	for i := range depthsInt[:len(depthsInt)-1] {
		if depthsInt[i] < depthsInt[i+1] {
			count++
		}
	}

	fmt.Printf("Count: %v\n", count)

	count = 0
	for i := range depthsInt[:len(depthsInt)-3] {
		if depthsInt[i] < depthsInt[i+3] {
			count++
		}
	}
	fmt.Printf("Sliding Window Count: %v\n", count)
}
