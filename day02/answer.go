package day02

import (
	"fmt"

	"advent/util"
)

func Answer() {
	instructions := util.ReadTokens("day02/input")

	horizontalPosition := 0
	depth := 0
	for i := 0; i < len(instructions); i += 2 {
		arg := util.ReadInt(instructions[i+1])
		switch instructions[i] {
		case "forward":
			horizontalPosition += arg
		case "down":
			depth += arg
		case "up":
			depth -= arg
		}
	}

	fmt.Printf("Answer (Part 1): %v\n", horizontalPosition*depth)

	// Part 2
	horizontalPosition = 0
	depth = 0
	aim := 0
	for i := 0; i < len(instructions); i += 2 {
		arg := util.ReadInt(instructions[i+1])
		switch instructions[i] {
		case "forward":
			horizontalPosition += arg
			depth += aim * arg
		case "down":
			aim += arg
		case "up":
			aim -= arg
		}
	}
	fmt.Printf("Answer (Part 2): %v\n", horizontalPosition*depth)

}
