package sumofprimeindexeselement

import (
	"math/big"
)

func Solve_s(arr []int) int {
	c := 0
	for i := 0; i < len(arr); i++ {
		n := arr[i]
		if big.NewInt(int64(i)).ProbablyPrime(0) {
			c += n
		}
	}

	return c
}
