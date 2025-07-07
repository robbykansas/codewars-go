package humanreadabletime

import "fmt"

func HumanReadableTime(seconds int) string {
	if seconds > 359999 {
		return "99:59:59"
	}

	s := seconds % 60

	if seconds < 3600 {
		m := seconds / 60
		return fmt.Sprintf("00:%02d:%02d", m, s)
	}

	h := seconds / 3600
	m := (seconds - h*3600) / 60

	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}
