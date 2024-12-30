func merge(nums1 []int, m int, nums2 []int, n int) {
	merged := append(nums1[:m], nums2...)
	sort.Ints(merged)
	copy(merged, nums1)

}