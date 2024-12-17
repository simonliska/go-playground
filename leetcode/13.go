func romanToInt(s string) int {
	list := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	x := 0
	pre := ""

	for _, c := range s {
		if list[string(c)] > list[pre] {
			x = x - list[pre]
			x = x + (list[string(c)] - list[pre])
		} else {
			x += list[string(c)]
		}
		pre = string(c)
	}
	return x
}