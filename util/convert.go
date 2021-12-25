package util

func ListToInt(items []string) (ints []int) {
	for _, item := range items {
		ints = append(ints, ReadInt(item))
	}
	return
}
