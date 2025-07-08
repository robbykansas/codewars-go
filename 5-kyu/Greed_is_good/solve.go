package greedisgood

func Score(dice [5]int) int {
	arr := make(map[int]int)
	score := 0

	for _, d := range dice {
		arr[d] += 1
	}

	for i, k := range arr {
		if k >= 3 {
			if i == 1 {
				score += 1000
			} else {
				score += (i * 100)
			}
			arr[i] -= 3
		}
	}

	if arr[1] >= 1 {
		score += (arr[1] * 100)
	}

	if arr[5] >= 1 {
		score += (arr[5] * 50)
	}

	return score
}
