package isbn10validation

import "strconv"

func ValidISBN10(isbn string) bool {
	if len(isbn) != 10 {
		return false
	}
	count := 0
	for i, d := range isbn {
		if i == 9 && (d == 'X' || d == 'x') {
			count += 100
		} else {
			n, err := strconv.Atoi(string(isbn[i]))
			if err != nil {
				return false
			}

			count += n * (i + 1)
		}
	}

	if count%11 == 0 {
		return true
	}
	return false // implement proper solution here
}
