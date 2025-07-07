package nextbiggernumberafterdigit

import (
	"fmt"
	"sort"
	"strconv"
)

func NextBigger(n int) int {
	d := strconv.Itoa(n)
	pivot := -1

	if len(d) == 1 {
		return -1
	}

	// find a pivot first
	for i := len(d) - 1; i >= 1; i-- {
		if int(d[i-1]) < int(d[i]) {
			pivot = i - 1
			break
		}
	}

	if pivot == -1 {
		return pivot
	}

	//get the next big number after pivot, if d[pivot] = 1, get number 2
	left := d[:pivot]
	right := d[pivot:]

	fmt.Println(n, pivot, right)
	p, r := PivotSorter(right)
	fmt.Println(p, r)

	res := left + p + r
	resInt, _ := strconv.Atoi(res)
	return resInt
}

func PivotSorter(s string) (string, string) {
	digit := []byte(s)
	pivot, _ := strconv.Atoi(string(s[0]))
	pv := 0
	index := 0
	for i := 1; i <= len(s)-1; i++ {
		indexNumber, _ := strconv.Atoi(string(s[i]))
		if indexNumber > pivot {
			if pv == 0 || indexNumber < pv {
				pv = indexNumber
				index = i
			}
		}
	}

	fmt.Println(pv, index)
	digit[0], digit[index] = digit[index], digit[0]

	if len(s) == 2 {
		return "", string(digit)
	}

	sl := digit[1:]
	sort.Slice(sl, func(i, j int) bool {
		return sl[i] < sl[j]
	})

	return string(digit[0]), string(sl)
}
