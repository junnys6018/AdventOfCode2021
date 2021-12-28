package main

import (
	"advent/day01"
	"advent/day02"
	"advent/day03"
	"advent/day04"
	"advent/day05"
	"advent/day06"
	"advent/day07"
	"advent/day08"
	"advent/day09"

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
	case 3:
		day03.Answer()
	case 4:
		day04.Answer()
	case 5:
		day05.Answer()
	case 6:
		day06.Answer()
	case 7:
		day07.Answer()
	case 8:
		day08.Answer()
	case 9:
		day09.Answer()
	default:
		fmt.Fprintf(os.Stderr, "Invalid day: %v\n", day)
	}
}
