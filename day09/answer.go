package day09

import (
	"advent/util"
	"fmt"
	"sort"
)

func bound(v, max int) bool {
	return v >= 0 && v < max
}

type Queue [][2]int

func (q *Queue) empty() bool {
	return len(*q) == 0
}

func (q *Queue) pop() [2]int {
	if q.empty() {
		panic("Attempt to pop empty queue")
	}
	top := (*q)[0]

	*q = (*q)[1:]

	return top
}

func (q *Queue) push(v [2]int) {
	*q = append(*q, v)
}

func basin(x, y int, floor [][]int) (size int) {
	height := len(floor)
	width := len(floor[0])
	visited := make([]bool, width*height)

	queue := Queue{}
	queue.push([2]int{x, y})

	deltas := [...][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	for !queue.empty() {
		c := queue.pop()
		good := true
		pending := [][2]int{}
		for _, delta := range deltas {
			dx := c[0] + delta[0]
			dy := c[1] + delta[1]
			if bound(dx, width) && bound(dy, height) && !visited[dy*width+dx] && floor[dy][dx] != 9 {
				pending = append(pending, [2]int{dx, dy})
				if floor[c[1]][c[0]] > floor[dy][dx] {
					good = false
				}
			}
		}

		if good {
			visited[c[1]*width+c[0]] = true
			for _, p := range pending {
				queue.push(p)
			}
		}
	}

	for _, v := range visited {
		if v {
			size++
		}
	}
	return size
}

func Answer() {
	floor := util.ReadGridInt("day09/input")
	height := len(floor)
	width := len(floor[0])

	deltas := [...][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	lowPoints := [][2]int{}
	risk := 0
	for y := range floor {
		for x := range floor[y] {
			min := true
			depth := floor[y][x]
			for _, delta := range deltas {
				dx := x + delta[0]
				dy := y + delta[1]
				if bound(dx, width) && bound(dy, height) && depth >= floor[dy][dx] {
					min = false
					break
				}
			}
			if min {
				lowPoints = append(lowPoints, [2]int{x, y})
				risk += 1 + depth
			}
		}
	}

	fmt.Printf("Answer (Part 1): %v\n", risk)

	basins := []int{}
	for _, p := range lowPoints {
		size := basin(p[0], p[1], floor)
		basins = append(basins, size)
	}
	sort.Ints(basins)
	fmt.Printf("Answer (Part 2): %v\n", basins[len(basins)-1]*basins[len(basins)-2]*basins[len(basins)-3])

}
