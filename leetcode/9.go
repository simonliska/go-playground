func isPalindrome(x int) bool {

	o := strconv.Itoa(x)
	r := ""

	for _, c := range o {
		r = string(c) + r
	}
	return o == r
}