package day05

import (
	"bufio"
	"fmt"
	"os"
)

type ChunkCoord struct {
	x uint
	y uint
}

type Map struct {
	// 16x16 chunks
	chunks map[ChunkCoord][]uint
}

func (m *Map) addChunk(x, y uint) {
	_, ok := m.chunks[ChunkCoord{x, y}]
	if !ok {
		m.chunks[ChunkCoord{x, y}] = make([]uint, 16*16)
	}
}

func (m *Map) paint(x, y uint) {
	xChunk := x / 16
	yChunk := y / 16

	xFine := x % 16
	yFine := y % 16

	m.addChunk(xChunk, yChunk)

	m.chunks[ChunkCoord{xChunk, yChunk}][yFine*16+xFine]++
}

func (m *Map) paintLine(x1, y1, x2, y2 uint) {
	if x1 == x2 {

		if y1 > y2 {
			y1, y2 = y2, y1
		}

		for y := y1; y <= y2; y++ {
			m.paint(x1, y)
		}

	} else if y1 == y2 {

		if x1 > x2 {
			x1, x2 = x2, x1
		}

		for x := x1; x <= x2; x++ {
			m.paint(x, y1)
		}

	} else {
		dir := func(a, b uint) int {
			if a > b {
				return -1
			}
			return 1
		}

		dx := dir(x1, x2)
		dy := dir(y1, y2)

		for x1 != x2 || y1 != y2 {
			m.paint(x1, y1)
			x1 = uint(int(x1) + dx)
			y1 = uint(int(y1) + dy)
		}
		m.paint(x1, y1)

	}
}

func (m *Map) count() (ret int) {
	for _, v := range m.chunks {
		for _, cnt := range v {
			if cnt >= 2 {
				ret++
			}
		}
	}
	return
}

func makeMap() Map {
	return Map{chunks: make(map[ChunkCoord][]uint)}
}

func Answer() {
	m1 := makeMap()
	m2 := makeMap()

	file, err := os.Open("day05/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var x1, y1, x2, y2 uint
		fmt.Sscanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if x1 == x2 || y1 == y2 {
			m1.paintLine(x1, y1, x2, y2)
		}
		m2.paintLine(x1, y1, x2, y2)

	}

	fmt.Printf("Answer (Part 1): %d\n", m1.count())
	fmt.Printf("Answer (Part 2): %d\n", m2.count())
}
