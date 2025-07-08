package countthedivisiblenumbers

func DivisibleCount(x, y, k uint64) uint64 {
	if y < k {
		return 1
	}
	return (y / k) - ((x - 1) / k)
}
