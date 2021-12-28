package day10

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func parse(path string) (ret []string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	return
}

func open(ch rune) bool {
	return ch == '(' || ch == '[' || ch == '{' || ch == '<'
}

func match(open, close rune) bool {
	return open == '(' && close == ')' ||
		open == '[' && close == ']' ||
		open == '{' && close == '}' ||
		open == '<' && close == '>'
}

func value(ch rune) int {
	switch ch {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137

	case '(':
		return 1
	case '[':
		return 2
	case '{':
		return 3
	case '<':
		return 4
	}
	return 0
}

func Answer() {
	lines := parse("day10/input")
	uncompletedLines := []string{}

	points := 0
	for _, line := range lines {
		stack := []rune{}
		good := true
		for _, ch := range line {
			if open(ch) {
				stack = append(stack, ch)
			} else {
				if match(stack[len(stack)-1], ch) {
					stack = stack[:len(stack)-1]
				} else {
					// syntax error
					// fmt.Printf("%v - Expected %c, but found %c instead.\n", line, stack[len(stack)-1], ch)
					points += value(ch)
					good = false
					break
				}
			}
		}

		if good {
			uncompletedLines = append(uncompletedLines, line)
		}
	}
	fmt.Printf("Answer (Part 1): %v\n", points)

	scores := []int{}
	for _, line := range uncompletedLines {
		stack := []rune{}
		for _, ch := range line {
			if open(ch) {
				stack = append(stack, ch)
			} else {
				if match(stack[len(stack)-1], ch) {
					stack = stack[:len(stack)-1]
				} else {
					panic("this is bad")
				}
			}
		}
		// pop the remaining stack
		score := 0
		for i := len(stack) - 1; i >= 0; i-- {
			score = score*5 + value(stack[i])
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)

	fmt.Printf("Answer (Part 1): %v\n", scores[(len(scores)-1)/2])

}
