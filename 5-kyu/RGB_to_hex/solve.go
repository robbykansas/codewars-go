package rgbtohex

import "fmt"

func RGB(r, g, b int) string {
	clamp := func(x int) int {
		if x < 0 {
			return 0
		}
		if x > 255 {
			return 255
		}
		return x
	}
	return fmt.Sprintf("%02X%02X%02X", clamp(r), clamp(g), clamp(b))
}
