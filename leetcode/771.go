func numJewelsInStones(jewels string, stones string) int {
	jewelSet := make(map[rune]bool)

	for _, jewel := range jewels {
		jewelSet[jewel] = true
	}

	count := 0
	for _, stone := range stones {
		if jewelSet[stone] {
			count++
		}
	}

	return count
}