package countingduplicates

import (
	"strings"
)

func duplicate_count(s1 string) int {
	s1 = strings.ToLower(s1)
	w := ""
	r := ""
	res := 0
	for _, d := range s1 {
		if !strings.Contains(w, string(d)) {
			w += string(d)
		} else {
			if !strings.Contains(r, string(d)) {
				r += string(d)
				res++
			}
		}
	}
	return res
}
