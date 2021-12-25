package day04

import (
	"advent/util"
	"fmt"
	"os"
	"strings"
)

type Board struct {
	numbers [25]int
	marked  [25]bool
}

func (b *Board) checkOff(n int) {
	for i, v := range b.numbers {
		if n == v {
			b.marked[i] = true
		}
	}
}

func (b *Board) winner() bool {
	// Check columns
	for x := 0; x < 5; x++ {
		win := true
		for y := 0; y < 5; y++ {
			win = win && b.marked[y*5+x]
		}
		if win {
			return true
		}
	}

	// Check rows
	for y := 0; y < 5; y++ {
		win := true
		for x := 0; x < 5; x++ {
			win = win && b.marked[y*5+x]
		}
		if win {
			return true
		}
	}
	return false
}

func (b *Board) score() (sum int) {
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if !b.marked[y*5+x] {
				sum += b.numbers[y*5+x]
			}
		}
	}
	return
}

func parse(path string) (numbers []int, boards []Board) {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	newLine := 0
	for i, v := range content {
		if v == '\n' {
			newLine = i
			break
		}
	}

	numbers = util.ListToInt(strings.Split(string(content[:newLine]), ","))

	boardNumbers := util.ListToInt(strings.Fields(string(content[newLine:])))

	for len(boardNumbers) > 0 {
		board := Board{}
		copy(board.numbers[:], boardNumbers[:25])
		boards = append(boards, board)

		boardNumbers = boardNumbers[25:]
	}

	return
}

func Answer() {
	numbers, boards := parse("day04/input")

	// Part 1
loop:
	for _, v := range numbers {
		for i := range boards {
			boards[i].checkOff(v)
			if boards[i].winner() {
				fmt.Printf("Answer (Part 1): %v\n", boards[i].score()*v)
				break loop
			}
		}
	}

	// Part 2
	winners := make(map[int]bool)
	numWinners := 0
loop2:
	for _, v := range numbers {
		for i := range boards {
			boards[i].checkOff(v)
			if boards[i].winner() {
				_, ok := winners[i]
				if !ok {
					winners[i] = true
					numWinners++
					if numWinners == len(boards) {
						fmt.Printf("Answer (Part 2): %v\n", boards[i].score()*v)
						break loop2
					}
				}
			}
		}
	}
}
