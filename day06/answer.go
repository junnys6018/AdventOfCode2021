package day06

import (
	"advent/util"
	"fmt"
	"os"
	"strings"
)

func Answer() {
	fish := make([]int, 9)

	content, err := os.ReadFile("day06/input")
	if err != nil {
		panic(err)
	}

	counts := util.ListToInt(strings.Split(string(content), ","))

	for _, v := range counts {
		fish[v]++
	}

	for i := 0; i < 80; i++ {
		numZeroDays := fish[0]

		for j := 0; j < 8; j++ {
			fish[j] = fish[j+1]
		}
		fish[8] = numZeroDays
		fish[6] += numZeroDays
	}

	sum := 0
	for _, v := range fish {
		sum += v
	}

	fmt.Printf("Answer (Part 1): %d\n", sum)
}
