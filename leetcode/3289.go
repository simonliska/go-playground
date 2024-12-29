func getSneakyNumbers(nums []int) []int {
	slices.Sort(nums)
	var output []int
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			output = append(output, nums[i])
		}
	}
	return output
}