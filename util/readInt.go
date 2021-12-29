package util

import (
	"bufio"
	"os"
	"strconv"
)

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

func ReadGridInt(path string) (ret [][]int) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		arr := []int{}
		for _, ch := range line {
			arr = append(arr, int(ch)-int('0'))
		}
		ret = append(ret, arr)
	}
	return
}
