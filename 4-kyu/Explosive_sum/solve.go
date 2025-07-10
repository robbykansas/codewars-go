package explosivesum

func ExpSum(n uint64) uint64 {
	memo := make([]uint64, n+1)
	memo[0] = 1

	for i := 1; i <= int(n); i++ {
		for j := i; j <= int(n); j++ {
			memo[j] += memo[j-i]
		}
	}

	return memo[n]
}
