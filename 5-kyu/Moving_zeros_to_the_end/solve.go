package movingzerostotheend

func MoveZeros(arr []int) []int {
	order := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[i], arr[order] = arr[order], arr[i]
			order++
		}
	}

	return arr
}
