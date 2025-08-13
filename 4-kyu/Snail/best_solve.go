package snail

func Snail_s(snailMap [][]int) []int {

	xmin := 0
	ymin := 0
	xmax := len(snailMap[0]) - 1
	ymax := len(snailMap) - 1
	resultln := len(snailMap[0]) * len(snailMap)
	result := make([]int, 0)

	for len(result) < resultln {

		for x := xmin; x <= xmax; x++ {
			result = append(result, snailMap[ymin][x])
		}
		ymin++

		for y := ymin; y <= ymax; y++ {
			result = append(result, snailMap[y][xmax])
		}
		xmax--

		for x := xmax; x >= xmin; x-- {
			result = append(result, snailMap[ymax][x])
		}
		ymax--

		for y := ymax; y >= ymin; y-- {
			result = append(result, snailMap[y][xmin])
		}
		xmin++
	}
	return result
}
