func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxValue := slices.Max(candies)
	var result []bool

	for _, x := range candies {
		if x+extraCandies >= maxValue {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}
	return result
}