package firstnonrepeatingcharacter

import (
	"strings"
)

func FirstNonRepeating_s(str string) string {
	for _, c := range str {
		if strings.Count(strings.ToLower(str), strings.ToLower(string(c))) < 2 {
			return string(c)
		}
	}
	return ""
}
