func scoreOfString(s string) (sum int) {
	for i := 0; i < len(s)-1; i++ {
		sum += abs(int(s[i]) - int(s[i+1]))
	}
	return sum
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}