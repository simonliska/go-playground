func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		target := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if target == nums[j] {
				return []int{i, j}
			}
		}
	}
	return nil
}