package day11

import (
	"advent/util"
	"fmt"
)

func canFlash(level [][]int, flashed [100]bool) bool {
	for i := range level {
		for j := range level[i] {
			if level[i][j] >= 10 && !flashed[i*10+j] {
				return true
			}
		}
	}
	return false
}

func step(level [][]int) (numFlashed int) {
	// Increment
	for i := range level {
		for j := range level[i] {
			level[i][j]++
		}
	}

	deltas := [...][2]int{
		{1, 1},
		{1, 0},
		{1, -1},
		{0, 1},
		{0, -1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
	}

	flashed := [10 * 10]bool{}

	for canFlash(level, flashed) {
		for i := range level {
			for j := range level[i] {
				if level[i][j] >= 10 && !flashed[i*10+j] {
					numFlashed++
					flashed[i*10+j] = true
					for _, delta := range deltas {
						di := i + delta[0]
						dj := j + delta[1]
						if di >= 0 && di < 10 && dj >= 0 && dj < 10 {
							level[di][dj]++
						}
					}
				}
			}
		}
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if flashed[i*10+j] {
				level[i][j] = 0
			}
		}
	}
	return
}

// func print(level [][]int) {
// 	fmt.Println("##### LEVEL #####")
// 	for i := range level {
// 		for j := range level[i] {
// 			fmt.Printf("%v", level[i][j])
// 		}
// 		fmt.Printf("\n")
// 	}
// }

func Answer() {
	level := util.ReadGridInt("day11/input")
	// print(level)
	// step(level)
	// print(level)

	sum := 0
	for i := 0; i < 100; i++ {
		sum += step(level)
	}

	fmt.Printf("Answer (Part 1): %v\n", sum)

	level = util.ReadGridInt("day11/input")
	for i := 0; ; i++ {
		if step(level) == 100 {
			fmt.Printf("Answer (Part 2): %v\n", i+1)
			break
		}
	}
}
