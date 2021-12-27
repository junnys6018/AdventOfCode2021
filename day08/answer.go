package day08

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type TestCase struct {
	inputs  [10]string
	outputs [4]string
}

func (tc *TestCase) countUniqueDigits() (count int) {
	for _, v := range tc.outputs {
		length := len(v)
		if length == 2 || length == 4 || length == 3 || length == 7 {
			count++
		}
	}
	return
}

func (tc *TestCase) solve() int {
	ord := func(ch rune) int {
		return int(ch) - int('a')
	}
	transform := func(s string) (ret int) {
		for _, ch := range s {
			ret |= 1 << ord(ch)
		}
		return
	}

	inputs := [10]int{}
	outputs := [4]int{}

	for i, s := range tc.inputs {
		inputs[i] = transform(s)
	}

	for i, s := range tc.outputs {
		outputs[i] = transform(s)
	}

	// maps an input to the digit it represents
	solution := map[int]int{}

	// maps a digit to the input that generates that digit
	invSolution := [10]int{}

	numActive := func(digit int) (ret int) {
		for i := 0; i < 7; i++ {
			if digit&(1<<i) != 0 {
				ret++
			}
		}
		return
	}

	// Solve 1, 7, 4 and 8
	for _, input := range inputs {
		switch numActive(input) {
		case 2:
			solution[input] = 1
			invSolution[1] = input
		case 3:
			solution[input] = 7
			invSolution[7] = input
		case 4:
			solution[input] = 4
			invSolution[4] = input
		case 7:
			solution[input] = 8
			invSolution[8] = input
		}
	}

	// solve for 3
	for _, input := range inputs {
		if numActive(input) == 5 && numActive(invSolution[1]&input) == 2 {
			solution[input] = 3
			invSolution[3] = input
			break
		}
	}

	// solve for 6
	for _, input := range inputs {
		if numActive(input) == 6 && numActive(invSolution[1]&input) == 1 {
			solution[input] = 6
			invSolution[6] = input
			break
		}
	}

	// solve for 5
	for _, input := range inputs {
		if numActive(input) == 5 && numActive(invSolution[6]&input) == 5 {
			solution[input] = 5
			invSolution[5] = input
			break
		}
	}

	// solve for 2
	for _, input := range inputs {
		if numActive(input) == 5 && solution[input] == 0 {
			solution[input] = 2
			invSolution[2] = input
			break
		}
	}

	// solve for 9
	for _, input := range inputs {
		if numActive(input) == 6 && numActive(invSolution[4]&input) == 4 {
			solution[input] = 9
			invSolution[9] = input
			break
		}
	}

	// for i, input := range inputs {
	// 	fmt.Println(tc.inputs[i], solution[input])
	// }

	answer := 0
	for digit, output := range outputs {
		answer += int(math.Pow10(3-digit)) * solution[output]
	}
	return answer
}

func Answer() {
	file, err := os.Open("day08/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var testCases []TestCase

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		var testCase TestCase

		copy(testCase.inputs[:], tokens[:10])
		copy(testCase.outputs[:], tokens[11:])
		testCases = append(testCases, testCase)
	}

	sum1 := 0
	sum2 := 0
	for i := range testCases {
		sum1 += testCases[i].countUniqueDigits()
		sum2 += testCases[i].solve()
	}

	fmt.Printf("Answer (Part 1): %v\n", sum1)
	fmt.Printf("Answer (Part 2): %v\n", sum2)

}
