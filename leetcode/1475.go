func finalPrices(prices []int) []int {
	var x []int
	for i := 0; i < len(prices); i++ {
		discount := 0
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				discount = prices[j]
				break
			}
		}
		x = append(x, prices[i]-discount)
	}
	return x
}