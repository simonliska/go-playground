func shuffle(nums []int, n int) []int {
	x := make([]int, 0, 2*n)

	for i := 0; i < n; i++ {
		x = append(x, nums[i])
		x = append(x, nums[i+n])
	}
	return x
}