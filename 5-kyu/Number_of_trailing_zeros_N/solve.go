package numberoftrailingzerosn

func Zeros(n int) int {
	res := 0
	for i := 5; n/i > 0; i *= 5 {
		res += n / i
	}
	return res
}
