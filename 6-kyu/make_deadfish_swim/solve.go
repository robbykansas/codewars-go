package makedeadfishswim

import "strings"

func Parse(data string) []int {
	var res []int
	var count int = 0

	if !strings.Contains(data, "o") {
		return []int{}
	}

	for i := range data {
		switch string(data[i]) {
		case "i":
			count += 1
		case "d":
			count -= 1
		case "s":
			count *= count
		case "o":
			res = append(res, count)
		}
	}

	return res
}
