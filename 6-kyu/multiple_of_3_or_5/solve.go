package multipleof3or5

func Multiple3And5(number int) int {
	res := 0
	for i := 0; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			res += i
		}
	}
	return res
}
