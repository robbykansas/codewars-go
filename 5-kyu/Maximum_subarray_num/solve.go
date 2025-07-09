package maximumsubarraynum

func MaximumSubarraySum(numbers []int) int {
	max := 0
	curr := 0
	for _, num := range numbers {
		curr = maxNum(num, curr+num)
		max = maxNum(max, curr)
	}

	return max
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
