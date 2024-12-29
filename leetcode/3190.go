func minimumOperations(nums []int) int {
	operations := 0
	for i := 0; i < len(nums); i++ {
		for nums[i]%3 != 0 {
			if nums[i]%3 == 1 {
				nums[i] -= 1
			} else if nums[i]%3 == 2 {
				nums[i] += 1
			}
			operations++
		}
	}
	return operations
}