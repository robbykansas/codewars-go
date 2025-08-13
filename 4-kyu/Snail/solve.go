package snail

func Snail(snaipMap [][]int) []int {
	ans := []int{}
	if len(snaipMap) == 0 {
		return []int{}
	}

	m, n := len(snaipMap), len(snaipMap[0])
	const (
		UP = iota
		RIGHT
		DOWN
		LEFT
	)

	direction := RIGHT
	dir_up, dir_right, dir_down, dir_left := 0, n, m, -1
	i, j := 0, 0

	for len(ans) != m*n {
		switch direction {
		case RIGHT:
			for j < dir_right {
				ans = append(ans, snaipMap[i][j])
				j++
			}
			i++
			j--
			dir_right--
			direction = DOWN
		case DOWN:
			for i < dir_down {
				ans = append(ans, snaipMap[i][j])
				i++
			}
			i--
			j--
			dir_down--
			direction = LEFT
		case LEFT:
			for j > dir_left {
				ans = append(ans, snaipMap[i][j])
				j--
			}
			i--
			j++
			dir_left++
			direction = UP
		default:
			for i > dir_up {
				ans = append(ans, snaipMap[i][j])
				i--
			}
			i++
			j++
			dir_up++
			direction = RIGHT
		}
	}

	return ans
}
