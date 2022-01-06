package day18

import (
	"advent/util"
	"errors"
	"fmt"
	"unicode"
)

type Number struct {
	value, depth        int
	left, right, parent *Number
}

func (n Number) IsLiteral() bool {
	return n.left == nil
}

func (n Number) String() string {
	if n.IsLiteral() {
		return fmt.Sprintf("%d", n.value)
	}

	return fmt.Sprintf("[%s,%s]", n.left, n.right)
}

func (n *Number) DeepClone() *Number {
	if n.IsLiteral() {
		return &Number{value: n.value, depth: n.depth}
	}

	left := n.left.DeepClone()
	right := n.right.DeepClone()

	num := &Number{depth: n.depth, left: left, right: right}

	left.parent = num
	right.parent = num

	return num
}

func Add(n1, n2 *Number) *Number {
	n1.IncreaseDepth()
	n2.IncreaseDepth()

	n := &Number{left: n1, right: n2}
	n1.parent = n
	n2.parent = n

	n.ReduceAll()

	return n
}

func (n *Number) IncreaseDepth() {
	n.depth++
	if !n.IsLiteral() {
		n.left.IncreaseDepth()
		n.right.IncreaseDepth()
	}
}

type ReduceResult int
type Side int

const (
	NONE ReduceResult = iota
	EXPLODE
	SPLIT

	LEFT Side = iota
	RIGHT
)

func (n *Number) GetChild(side Side) *Number {
	switch side {
	case LEFT:
		return n.left
	case RIGHT:
		return n.right
	default:
		panic("bad side")
	}
}

func (n *Number) Magnitude() int {
	if n.IsLiteral() {
		return n.value
	}
	return 3*n.left.Magnitude() + 2*n.right.Magnitude()
}

// Finds the next left and right literal in the tree
func (n *Number) FindLiterals() (*Number, *Number) {
	parent := n.parent

	var s1, s2 Side

	if parent.right == n {
		s1 = LEFT
		s2 = RIGHT
	} else {
		s1 = RIGHT
		s2 = LEFT
	}

	s1Literal := parent.GetChild(s1)
	for !s1Literal.IsLiteral() {
		s1Literal = s1Literal.GetChild(s2)
	}

	ansestor := parent.parent
	for ansestor.GetChild(s2) == parent {
		parent = ansestor
		ansestor = ansestor.parent
		if ansestor == nil {
			break
		}
	}

	if ansestor == nil {
		if s1 == LEFT {
			return s1Literal, nil
		}
		return nil, s1Literal
	}

	s2Literal := ansestor.GetChild(s2)
	for !s2Literal.IsLiteral() {
		s2Literal = s2Literal.GetChild(s1)
	}

	if s1 == LEFT {
		return s1Literal, s2Literal
	}
	return s2Literal, s1Literal
}

func (n *Number) ReduceExplode() ReduceResult {
	if !n.IsLiteral() && n.depth == 4 /* explode */ {
		leftLiteral, rightLiteral := n.FindLiterals()

		if !n.left.IsLiteral() || !n.right.IsLiteral() {
			panic("bad number")
		}

		if leftLiteral != nil {
			leftLiteral.value += n.left.value
		}

		if rightLiteral != nil {
			rightLiteral.value += n.right.value
		}

		n.left = nil
		n.right = nil
		n.value = 0

		return EXPLODE
	}

	if !n.IsLiteral() {
		result := n.left.ReduceExplode()
		if result != NONE {
			return result
		}
		result = n.right.ReduceExplode()
		return result
	}

	return NONE
}

func (n *Number) ReduceSplit() ReduceResult {
	if n.IsLiteral() && n.value > 9 /* split */ {
		n1 := &Number{value: n.value / 2, depth: n.depth + 1, parent: n}
		n2 := &Number{value: (n.value + 1) / 2, depth: n.depth + 1, parent: n}

		n.value = 0 // for saftey
		n.left = n1
		n.right = n2

		return SPLIT
	}

	if !n.IsLiteral() {
		result := n.left.ReduceSplit()
		if result != NONE {
			return result
		}
		result = n.right.ReduceSplit()
		return result
	}

	return NONE
}

func (n *Number) Reduce() ReduceResult {
	result := n.ReduceExplode()
	if result == NONE {
		result = n.ReduceSplit()
	}

	return result
}

func (n *Number) ReduceAll() {
	for n.Reduce() != NONE {
	}
}

func NewNumberString(s string, depth int) (*Number, int, error) {
	current := 0

	token := s[current]
	current++

	if token == '[' {
		n1, read, err := NewNumberString(s[current:], depth+1)
		current += read

		if err != nil {
			return nil, 0, errors.New("bad string")
		}

		// expect a comma
		if s[current] != ',' {
			return nil, 0, errors.New("bad string")
		}
		current++

		n2, read, err := NewNumberString(s[current:], depth+1)
		current += read
		if err != nil {
			return nil, 0, errors.New("bad string")
		}

		// expect closing paren
		if s[current] != ']' {
			return nil, 0, errors.New("bad string")
		}
		current++

		n := &Number{left: n1, right: n2, depth: depth}
		n1.parent = n
		n2.parent = n

		return n, current, nil

	} else if unicode.IsDigit(rune(token)) {
		return &Number{value: int(token - '0'), depth: depth}, current, nil
	} else {
		return nil, 0, errors.New("bad string")
	}
}

func part1(input string) {
	lines := util.ReadLines(input)

	num, _, err := NewNumberString(lines[0], 0)
	if err != nil {
		panic(err)
	}

	for _, s := range lines[1:] {
		nextNum, _, err := NewNumberString(s, 0)
		if err != nil {
			panic(err)
		}

		num = Add(num, nextNum)
	}

	fmt.Printf("Answer (Part 1): %v\n", num.Magnitude())
}

func part2(input string) {
	lines := util.ReadLines(input)

	nums := []*Number{}

	for _, s := range lines {
		num, _, err := NewNumberString(s, 0)
		if err != nil {
			panic(err)
		}

		nums = append(nums, num)
	}

	max := 0
	for _, n1 := range nums {
		for _, n2 := range nums {
			if n1 == n2 {
				continue
			}

			add := Add(n1.DeepClone(), n2.DeepClone())
			mag := add.Magnitude()
			if mag > max {
				max = mag
			}
		}
	}

	fmt.Printf("Answer (Part 2): %v\n", max)
}

func Answer() {
	input := "day18/input"
	part1(input)
	part2(input)
}
