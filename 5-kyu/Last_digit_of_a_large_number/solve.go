package lastdigitofalargenumber

import "strconv"

func LastDigit(n1, n2 string) int {
	d := digit(n1, n2)

	return d
}

func digit(n1, n2 string) int {
	lastDigit := string(n1[len(n1)-1])

	if lastDigit == "0" || lastDigit == "1" || lastDigit == "5" || lastDigit == "6" {
		d, _ := strconv.Atoi(lastDigit)
		return d
	}

	digit := make(map[string][]string)
	digit["2"] = []string{"2", "4", "8", "6"}
	digit["3"] = []string{"3", "9", "7", "1"}
	digit["4"] = []string{"4", "6"}
	digit["7"] = []string{"7", "9", "3", "1"}
	digit["8"] = []string{"8", "4", "2", "6"}
	digit["9"] = []string{"9", "1"}

	cycle := digit[lastDigit]
	n := modString(n2, len(cycle))
	res := n % len(cycle)
	if res == 0 {
		res = len(cycle) - 1
	} else {
		res = res - 1
	}

	d, _ := strconv.Atoi(cycle[res])

	return d
}

func modString(nStr string, mod int) int {
	result := 0
	for _, ch := range nStr {
		digit := int(ch - '0')
		result = (result*10 + digit) % mod
	}
	return result
}
