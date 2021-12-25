package day03

import (
	"advent/util"
	"fmt"
	"strconv"
)

func count(x uint64) int {
	if x == 0 {
		return 0
	}
	return 1
}

func filter(items []uint64, predicate func(uint64) bool) (ret []uint64) {
	for _, v := range items {
		if predicate(v) {
			ret = append(ret, v)
		}
	}
	return
}

func Answer() {
	bitsStr := util.ReadTokens("day03/input")

	bitsInt := make([]uint64, len(bitsStr))
	for i := range bitsStr {
		bitsInt[i] = util.ReadBinary(bitsStr[i])
	}

	bitLength := len(bitsStr[0])
	counts := make([]int, bitLength)

	for _, v := range bitsInt {
		for i := 0; i < bitLength; i++ {
			counts[i] += count(v & (1 << i))
		}
	}

	gammaRate := uint64(0)
	for i := 0; i < bitLength; i++ {
		if counts[i] > len(bitsStr)/2 {
			gammaRate |= 1 << i
		}
	}

	epsilonRate := ^gammaRate & (1<<bitLength - 1)

	fmt.Printf(
		"Answer (Part 1): %v gamma=%v epsilon=%v\n",
		gammaRate*epsilonRate,
		strconv.FormatUint(gammaRate, 2),
		strconv.FormatUint(epsilonRate, 2),
	)

	// Part 2
	oxygenGeneratorRating := bitsInt

	for i := bitLength - 1; i >= 0; i-- {
		num := len(oxygenGeneratorRating) + 1

		counti := 0
		for _, v := range oxygenGeneratorRating {
			counti += count(v & (1 << i))
		}

		oxygenGeneratorRating = filter(oxygenGeneratorRating, func(item uint64) bool {
			if (counti >= num/2 && (item&(1<<i)) != 0) || (counti < num/2 && (item&(1<<i)) == 0) {
				return true
			}
			return false
		})

		fmt.Println(oxygenGeneratorRating, counti, num)

		if len(oxygenGeneratorRating) == 1 {
			break
		}
	}

	co2ScrubberRating := bitsInt

	for i := bitLength - 1; i >= 0; i-- {
		num := len(co2ScrubberRating) + 1

		counti := 0
		for _, v := range co2ScrubberRating {
			counti += count(v & (1 << i))
		}

		co2ScrubberRating = filter(co2ScrubberRating, func(item uint64) bool {
			if (counti >= num/2 && (item&(1<<i)) == 0) || (counti < num/2 && (item&(1<<i)) != 0) {
				return true
			}
			return false
		})

		fmt.Println(co2ScrubberRating, counti, num)

		if len(co2ScrubberRating) == 1 {
			break
		}
	}

	fmt.Printf(
		"Answer (Part 2): %v oxygen=%v c02=%v\n",
		oxygenGeneratorRating[0]*co2ScrubberRating[0],
		oxygenGeneratorRating[0],
		co2ScrubberRating[0],
	)
}
