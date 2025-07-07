package firstnonrepeatingcharacter

import (
	//   "fmt"
	"strings"
)

func FirstNonRepeating(str string) string {
	if len(str) == 1 {
		return str
	}

	ar := make(map[string]int)

	char := strings.Split(str, "")
	for _, c := range char {
		ar[strings.ToLower(c)] += 1
	}

	for _, v := range char {
		if ar[strings.ToLower(v)] == 1 {
			return v
		}
	}
	return ""
}
