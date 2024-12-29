func numIdenticalPairs(nums []int) int {
	pair := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && i < j {
				pair = pair + 1
			}
		}
	}
	return pair
}