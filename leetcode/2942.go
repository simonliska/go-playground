func findWordsContaining(words []string, x byte) []int {
	var result []int

	for i, word := range words {
		if strings.ContainsRune(word, rune(x)) {
			result = append(result, i)
		}
	}
	return result
}