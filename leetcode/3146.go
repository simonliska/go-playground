func findPermutationDifference(s string, t string) int {
	mapData := make(map[rune]int)
	for i, x := range s {
		mapData[x] = i
	}

	sum := 0
	for i, x := range t {
		sum += abs(mapData[x] - i)
	}
	return sum
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}