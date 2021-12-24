package main

import (
	"advent/day01"
	"advent/day02"

	"fmt"
	"os"
	"strconv"
)

func main() {
	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	switch day {
	case 1:
		day01.Answer()
	case 2:
		day02.Answer()
	default:
		fmt.Fprintf(os.Stderr, "Invalid day: %v\n", day)
	}
}
