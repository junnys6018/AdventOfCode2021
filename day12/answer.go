package day12

import (
	"advent/util"
	"fmt"
	"strings"
)

type Node struct {
	large    bool
	adjacent []*Node
	name     string
}

func isLarge(s string) bool {
	return strings.ToUpper(s) == s
}

func search(path []*Node, doubleVisit, part2 bool) (numPaths int) {
	end := path[len(path)-1]
	for _, next := range end.adjacent {
		if next.name == "start" {
			continue
		} else if next.name == "end" {
			numPaths++
			// for _, node := range path {
			// 	fmt.Printf("%v,", node.name)
			// }
			// fmt.Println("end")
		} else if next.large {
			newPath := make([]*Node, len(path)+1)
			copy(newPath, path)
			newPath[len(path)] = next

			numPaths += search(newPath, doubleVisit, part2)
		} else {
			found := false
			for _, visited := range path {
				if visited.name == next.name {
					found = true
					break
				}
			}

			if !found {
				newPath := make([]*Node, len(path)+1)
				copy(newPath, path)
				newPath[len(path)] = next

				numPaths += search(newPath, doubleVisit, part2)
			} else if part2 && !doubleVisit {
				newPath := make([]*Node, len(path)+1)
				copy(newPath, path)
				newPath[len(path)] = next

				numPaths += search(newPath, true, part2)
			}
		}
	}
	return
}

func Answer() {
	cave := make(map[string]*Node)

	links := util.ReadLines("day12/input")
	for _, link := range links {
		idx := strings.Index(link, "-")
		a := string(link[:idx])
		b := string(link[idx+1:])

		nodeA, ok := cave[a]
		if !ok {
			cave[a] = new(Node)
			nodeA = cave[a]
			nodeA.large = isLarge(a)
			nodeA.name = a
		}
		nodeB, ok := cave[b]
		if !ok {
			cave[b] = new(Node)
			nodeB = cave[b]
			nodeB.large = isLarge(b)
			nodeB.name = b
		}

		nodeA.adjacent = append(nodeA.adjacent, nodeB)
		nodeB.adjacent = append(nodeB.adjacent, nodeA)
	}
	fmt.Printf("Answer (Part 1): %v\n", search([]*Node{cave["start"]}, false, false))
	fmt.Printf("Answer (Part 2): %v\n", search([]*Node{cave["start"]}, false, true))
}
