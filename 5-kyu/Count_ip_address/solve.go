package countipaddress

import (
	"strconv"
	"strings"
)

func IpsBetween(start, end string) int {
	return ipInt(end) - ipInt(start)
}

func ipInt(ip string) int {
	s := strings.Split(ip, ".")
	count := 0
	for _, d := range s {
		parts, _ := strconv.Atoi(d)

		count = count*256 + parts
	}

	return count
}
