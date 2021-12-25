package day07

import (
	"advent/util"
	"fmt"
	"math"
)

func absInt(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func triangle(n int) int {
	return n * (n + 1) / 2
}

func Answer() {
	positions := util.ReadCommaSeperated("day07/input")
	fuel1 := math.MaxInt
	fuel2 := math.MaxInt

	maxPos := util.MaxInt(positions)
	for i := 0; i <= maxPos; i++ {
		sum1 := 0
		sum2 := 0
		for _, v := range positions {
			sum1 += absInt(v - i)
			sum2 += triangle(absInt(v - i))
		}
		if sum1 < fuel1 {
			fuel1 = sum1
		}
		if sum2 < fuel2 {
			fuel2 = sum2
		}
	}

	fmt.Printf("Answer (Part 1): %v\n", fuel1)
	fmt.Printf("Answer (Part 2): %v\n", fuel2)
}
