package solution

func kidsWithCandies(candies []int, extraCandies int) []bool {
	greaterNumber := 0
	result := []bool{}
	for _, candy := range candies {
		if candy > greaterNumber {
			greaterNumber = candy
		}
	}

	for _, candy := range candies {
		if candy+extraCandies >= greaterNumber {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}
	return result
}
