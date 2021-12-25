package util

import "math"

func SumInt(arr []int) (ret int) {
	for _, v := range arr {
		ret += v
	}
	return
}

func MaxInt(arr []int) int {
	max := math.MinInt
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}
