func countConsistentStrings(allowed string, words []string) int {
	result := 0

	allowedSet := make(map[rune]bool)
	for _, char := range allowed {
		allowedSet[char] = true
	}

	for _, word := range words {
		isConsistent := true
		for _, char := range word {
			if !allowedSet[char] {
				isConsistent = false
				break
			}
		}
		if isConsistent {
			result++
		}
	}

	return result
}