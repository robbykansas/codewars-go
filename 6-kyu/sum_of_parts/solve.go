package sumofparts

func PartsSums(ls []uint64) []uint64 {
	if len(ls) == 0 {
		return []uint64{0}
	}

	var max uint64
	res := []uint64{}

	for _, d := range ls {
		max += d
	}

	for i := 0; i <= len(ls)-1; i++ {
		if i == 0 {
			res = append(res, max)
		} else {
			max -= ls[i-1]
			res = append(res, max)
		}
	}

	res = append(res, 0)
	return res
}
