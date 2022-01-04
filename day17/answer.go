// Found a very sus analytical solution for this problem
// An explanation for how this works is left as an exercise to the reader
package day17

import (
	"fmt"
	"math"
	"os"
)

// Assumes that there exist an x such that x^2-sgn(x)x(x-1)/2 is in the range [x1, x2]
func part1(y1, y2 int) (yMax int) {
	for t := 1.0; t < 100000.0; t++ {
		lower := float64(y1)/t + (t-1.0)/2.0
		upper := float64(y2)/t + (t-1.0)/2.0

		if int(math.Floor(lower)) != int(math.Floor(upper)) || math.Floor(lower) == lower {
			yMax = int(math.Floor(upper))
		}
	}

	return yMax*(yMax+1) - (yMax+1)*yMax/2
}

func sgn(x int) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}

func part2(x1, x2, y1, y2 int) int {
	pairs := make(map[[2]int]bool)

	for t := 1.0; t < 100000.0; t++ {
		lower := float64(y1)/t + (t-1.0)/2.0
		upper := float64(y2)/t + (t-1.0)/2.0

		for x := 1; x < 100; x++ {
			fx := 0
			tInt := int(t)
			if tInt <= x {
				fx = tInt*x - sgn(x)*tInt*(tInt-1)/2
			} else {
				fx = x*x - sgn(x)*x*(x-1)/2
			}

			if fx >= x1 && fx <= x2 {
				for y := int(math.Ceil(lower)); y <= int(math.Floor(upper)); y++ {
					pairs[[2]int{x, y}] = true
				}
			}
		}
	}

	return len(pairs)
}

func Answer() {
	var x1, x2, y1, y2 int
	input, err := os.ReadFile("day17/input")
	if err != nil {
		panic(err)
	}

	fmt.Sscanf(string(input), "target area: x=%d..%d, y=%d..%d", &x1, &x2, &y1, &y2)

	fmt.Printf("Answer (Part 1): %v\n", part1(y1, y2))
	fmt.Printf("Answer (Part 2): %v\n", part2(x1, x2, y1, y2))
}
