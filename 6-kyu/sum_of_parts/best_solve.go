package sumofparts

func PartsSums_s(ls []uint64) []uint64 {
	v := make([]uint64, len(ls)+1)
	for i := len(ls) - 1; i >= 0; i-- {
		v[i] = v[i+1] + ls[i]
	}
	return v
}
