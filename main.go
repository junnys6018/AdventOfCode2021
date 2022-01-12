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
	"advent/day10"
	"advent/day11"
	"advent/day12"
	"advent/day13"
	"advent/day14"
	"advent/day15"
	"advent/day16"
	"advent/day17"
	"advent/day18"
	"advent/day19"
	"advent/day20"
	"advent/day21"

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
	case 10:
		day10.Answer()
	case 11:
		day11.Answer()
	case 12:
		day12.Answer()
	case 13:
		day13.Answer()
	case 14:
		day14.Answer()
	case 15:
		day15.Answer()
	case 16:
		day16.Answer()
	case 17:
		day17.Answer()
	case 18:
		day18.Answer()
	case 19:
		day19.Answer()
	case 20:
		day20.Answer()
	case 21:
		day21.Answer()
	default:
		fmt.Fprintf(os.Stderr, "Invalid day: %v\n", day)
	}
}
