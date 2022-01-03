package day15

import (
	"advent/util"
	"container/heap"
	"fmt"
	"math"
	"strings"
)

type Node struct {
	x, y, dist, index int
}

func (n Node) String() string {
	return fmt.Sprintf("{%d, %d, %d}", n.x, n.y, n.dist)
}

type Level [][]int

func (level Level) at(v [2]int) int {
	return level[v[1]][v[0]]
}

func (level Level) String() string {
	builder := strings.Builder{}

	for y := range level {
		for x := range level[y] {
			fmt.Fprintf(&builder, "%d", level[y][x])
		}
		if y != len(level)-1 {
			builder.WriteRune('\n')
		}
	}

	return builder.String()
}

// implements heap.interface
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Node)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	item := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return item
}

func dijkstra(level Level) int {
	height := len(level)
	width := len(level[0])

	visited := make(map[[2]int]struct{})
	nodes := make(map[[2]int]*Node)

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	for y := range level {
		for x := range level[y] {
			nodes[[2]int{x, y}] = &Node{x: x, y: y, dist: math.MaxInt}
			heap.Push(&pq, nodes[[2]int{x, y}])
		}
	}

	nodes[[2]int{0, 0}].dist = 0
	heap.Fix(&pq, nodes[[2]int{0, 0}].index)
	visited[[2]int{0, 0}] = struct{}{}

	deltas := [4][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)

		for _, delta := range deltas {
			v := [2]int{node.x + delta[0], node.y + delta[1]}

			if _, ok := visited[v]; !ok && v[0] >= 0 && v[0] < width && v[1] >= 0 && v[1] < height {
				newDist := node.dist + level.at(v)
				if newDist < nodes[v].dist {
					nodes[v].dist = newDist
					heap.Fix(&pq, nodes[v].index)
				}
			}
		}
	}

	return nodes[[2]int{width - 1, height - 1}].dist
}

func Answer() {
	level := Level(util.ReadGridInt("day15/input"))

	fmt.Printf("Answer (Part 1): %v\n", dijkstra(level))

	bigLevel := Level{}
	height := len(level)
	width := len(level[0])

	for y := 0; y < 5*height; y++ {
		line := make([]int, 5*width)
		for x := 0; x < 5*width; x++ {
			increment := x/width + y/height

			risk := level[y%height][x%width] + increment

			if risk > 9 {
				risk -= 9
			}

			line[x] = risk
		}
		bigLevel = append(bigLevel, line)
	}

	fmt.Printf("Answer (Part 2): %v\n", dijkstra(bigLevel))
}
