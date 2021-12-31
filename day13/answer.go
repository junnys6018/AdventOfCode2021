package day13

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

const (
	VERTICAL = iota
	HORIZONTAL
)

type Instruction struct {
	direction int
	v         int
}

func parse(path string) ([]Point, []Instruction) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	points := []Point{}
	instructions := []Instruction{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "fold") {
			var dir rune
			var a int
			fmt.Sscanf(line, "fold along %c=%d", &dir, &a)
			if dir == 'x' {
				instructions = append(instructions, Instruction{VERTICAL, a})
			} else if dir == 'y' {
				instructions = append(instructions, Instruction{HORIZONTAL, a})
			}
		} else {
			p := Point{}
			fmt.Sscanf(line, "%d,%d", &p.x, &p.y)
			points = append(points, p)
		}
	}

	return points, instructions
}

func removeDuplicates(points []Point) (newPoints []Point) {
	allKeys := make(map[Point]bool)
	for _, point := range points {
		if _, ok := allKeys[point]; !ok {
			allKeys[point] = true
			newPoints = append(newPoints, point)
		}
	}
	return
}

func fold(points []Point, instruction Instruction) (newPoints []Point) {
	if instruction.direction == HORIZONTAL {
		for _, point := range points {
			if point.y > instruction.v {
				newPoints = append(newPoints, Point{x: point.x, y: 2*instruction.v - point.y})
			} else {
				newPoints = append(newPoints, point)
			}
		}
	} else if instruction.direction == VERTICAL {
		for _, point := range points {
			if point.x > instruction.v {
				newPoints = append(newPoints, Point{x: 2*instruction.v - point.x, y: point.y})
			} else {
				newPoints = append(newPoints, point)
			}
		}
	}

	newPoints = removeDuplicates(newPoints)

	return
}

func printPoints(points []Point) {
	xMax := 0
	yMax := 0

	// Convert to set for quicker lookup
	pointSet := make(map[Point]bool)

	for _, point := range points {
		if point.x > xMax {
			xMax = point.x
		}
		if point.y > yMax {
			yMax = point.y
		}
		pointSet[point] = true
	}

	for y := 0; y <= yMax; y++ {
		for x := 0; x <= xMax; x++ {
			_, ok := pointSet[Point{x, y}]
			if ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}

}

func Answer() {
	points, instructions := parse("day13/input")

	points = fold(points, instructions[0])

	fmt.Printf("Answer (Part 1): %v\n", len(points))

	for _, instruction := range instructions[1:] {
		points = fold(points, instruction)
	}

	fmt.Println("Answer (Part 2):")
	printPoints(points)
}
