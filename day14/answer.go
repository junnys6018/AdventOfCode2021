package day14

import (
	"bufio"
	"fmt"
	"os"
)

type Rule struct {
	left, right, insert rune
}

type Pair struct {
	left, right rune
}

func parse(path string) (rune, rune, map[Pair]int, []Rule) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rules := []Rule{}

	scanner.Scan()
	template := scanner.Text()
	pairs := make(map[Pair]int)

	for i := range template[:len(template)-1] {
		pairs[Pair{rune(template[i]), rune(template[i+1])}]++
	}

	for scanner.Scan() {
		line := scanner.Text()
		rule := Rule{}
		fmt.Sscanf(line, "%c%c -> %c", &rule.left, &rule.right, &rule.insert)
		rules = append(rules, rule)
	}

	return rune(template[0]), rune(template[len(template)-1]), pairs, rules
}

func step(polymer map[Pair]int, rules []Rule) map[Pair]int {
	newPolymer := make(map[Pair]int)

	for pair, count := range polymer {
		replace := false
		for _, rule := range rules {
			if rule.left == pair.left && rule.right == pair.right {
				newPolymer[Pair{rule.left, rule.insert}] += count
				newPolymer[Pair{rule.insert, rule.right}] += count

				replace = true
				break
			}
		}

		if !replace {
			newPolymer[pair] += count
		}
	}

	return newPolymer
}

func answerPart(steps int) int {
	left, right, pairs, rules := parse("day14/input")

	for i := 0; i < steps; i++ {
		pairs = step(pairs, rules)
	}

	counts := make(map[rune]int)

	for pair, count := range pairs {
		counts[pair.left] += count
		counts[pair.right] += count
	}

	// Divide by two because each element is counted in 2 pairs
	for ch, count := range counts {
		counts[ch] = count / 2
	}

	// ... except for the left and rightmost elements
	counts[left]++
	counts[right]++

	min := left
	max := left

	for ch, count := range counts {
		if count > counts[max] {
			max = ch
		}
		if count < counts[min] {
			min = ch
		}
	}

	return counts[max] - counts[min]
}

func Answer() {
	fmt.Printf("Answer (Part 1): %v\n", answerPart(10))
	fmt.Printf("Answer (Part 2): %v\n", answerPart(40))
}
