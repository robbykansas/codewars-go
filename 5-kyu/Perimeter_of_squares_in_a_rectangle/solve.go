package perimeterofsquaresinarectangle

func Perimeter(n int) int {
	return 4 * fib(n+1)
}

func fib(n int) int {
	arr := []int{0, 1}
	count := 1
	for i := 2; i <= n; i++ {
		num := arr[i-2] + arr[i-1]
		arr = append(arr, num)
		count += num
	}

	return count
}
