package humanreadabledurationformat

import (
	"strconv"
	"strings"
)

func FormatDuration_s(seconds int64) string {
	if seconds == 0 {
		return "now"
	}
	m, seconds := seconds/60, seconds%60
	h, m := m/60, m%60
	d, h := h/24, h%24
	y, d := d/365, d%365
	q := queue(y, "year", nil)
	q = queue(d, "day", q)
	q = queue(h, "hour", q)
	q = queue(m, "minute", q)
	return join(queue(seconds, "second", q))
}
func queue(n int64, s string, q []string) []string {
	if n == 1 {
		return append(q, "1 "+s)
	}
	if n > 1 {
		return append(q, strconv.FormatInt(n, 10)+" "+s+"s")
	}
	return q
}
func join(q []string) string {
	if len(q) == 0 {
		return "now"
	}
	if len(q) == 1 {
		return q[0]
	}
	return strings.Join(append(q[:len(q)-2], q[len(q)-2]+" and "+q[len(q)-1]), ", ")
}
