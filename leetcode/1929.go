func getConcatenation(nums []int) []int {
	ans := nums
	return slices.Concat(nums, ans)
}