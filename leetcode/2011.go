func finalValueAfterOperations(operations []string) int {
	list := map[string]int{
		"++X": 1,
		"X++": 1,
		"--X": -1,
		"X--": -1,
	}
	x := 0

	for _, c := range operations {
		x = x + list[string(c)]
	}
	return x
}