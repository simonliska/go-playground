func maximumWealth(accounts [][]int) int {
	var result []int
	for i := 0; i < len(accounts); i++ {
		x := accounts[i]
		mid := 0
		for j := 0; j < len(x); j++ {
			mid += accounts[i][j]
		}
		result = append(result, mid)
	}
	return slices.Max(result)
}