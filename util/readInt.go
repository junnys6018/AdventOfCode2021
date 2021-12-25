package util

import "strconv"

func ReadInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return val
}

func ReadBinary(s string) uint64 {
	val, err := strconv.ParseUint(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return val
}
