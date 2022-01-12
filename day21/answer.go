package day21

import (
	"fmt"
	"os"
)

func parse(path string) [2]int {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	players := [2]int{}
	fmt.Sscanf(string(input), "Player 1 starting position: %d\nPlayer 2 starting position: %d", &players[0], &players[1])

	return players
}

func part1(players [2]int) {
	scores := [2]int{}

	dice := 1
	turn := 0

	for scores[0] < 1000 && scores[1] < 1000 {
		roll := 3*dice + 3
		dice += 3
		players[turn] = ((players[turn] + roll - 1) % 10) + 1
		scores[turn] += players[turn]

		turn = 1 - turn
	}

	loser := 0
	if scores[1] < 1000 {
		loser = 1
	}

	fmt.Printf("Answer (Part 1): %v\n", scores[loser]*(dice-1))
}

type GameState struct {
	players [2]int
	scores  [2]int
}

type Winner int

const (
	WINNER_P1   Winner = 0
	WINNER_P2   Winner = 1
	WINNER_NONE Winner = 2
)

func (gs GameState) winner() Winner {
	if gs.scores[0] >= 21 {
		return WINNER_P1
	} else if gs.scores[1] >= 21 {
		return WINNER_P2
	}
	return WINNER_NONE
}

func part2(players [2]int) {
	multiverse := make(map[GameState]uint64)
	multiverse[GameState{players, [2]int{}}] = 1

	branchLut := [...]int{1, 3, 6, 7, 6, 3, 1}
	turn := 0
	numWins := [2]uint64{}

	for len(multiverse) != 0 {

		nextMultiverse := make(map[GameState]uint64)

		for state, n := range multiverse {
			if state.winner() != WINNER_NONE {
				numWins[state.winner()] += n
				continue
			}

			for roll := 3; roll <= 9; roll++ {
				newState := state
				newState.players[turn] = ((newState.players[turn] + roll - 1) % 10) + 1
				newState.scores[turn] += newState.players[turn]

				nextMultiverse[newState] += n * uint64(branchLut[roll-3])
			}
		}

		turn = 1 - turn
		multiverse = nextMultiverse
	}

	winner := 0
	if numWins[1] > numWins[0] {
		winner = 1
	}

	fmt.Printf("Answer (Part 2): %v\n", numWins[winner])
}

func Answer() {
	players := parse("day21/input")
	part1(players)
	part2(players)
}
