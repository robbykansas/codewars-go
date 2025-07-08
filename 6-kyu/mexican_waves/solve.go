package mexicanwaves

import (
	"unicode"
)

func wave(words string) []string {
	// Your code here and happy coding!
	var res []string
	for i := 0; i < len(words); i++ {
		if string(words[i]) != " " {
			r := []rune(words)
			r[i] = unicode.ToUpper(r[i])
			res = append(res, string(r))
		}
	}

	return res
}
