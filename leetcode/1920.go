func buildArray(nums []int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		y := nums[nums[i]]
		result = append(result, y)
	}
	return result
}