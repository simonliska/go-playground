func smallestEvenMultiple(n int) int {
	for n%2 != 0 {
		n *= 2
	}
	return n
}