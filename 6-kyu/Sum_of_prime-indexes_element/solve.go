package sumofprimeindexeselement

func Solve(arr []int) int {
	sum := 0
	for i, num := range arr {
		if isPrime(i) {
			sum += num
		}
	}
	return sum
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
