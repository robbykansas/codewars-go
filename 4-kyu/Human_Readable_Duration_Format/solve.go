package humanreadabledurationformat

import "fmt"

func FormatDuration(seconds int64) string {
	var y_s, d_s, h_s, m_s, s_s string
	var convert func(y, d, h, m, s int64) string

	convert = func(y, d, h, m, s int64) string {
		arr := []string{}
		if y == 1 {
			y_s = fmt.Sprintf("%d year", y)
			arr = append(arr, y_s)
		} else if y > 1 {
			y_s = fmt.Sprintf("%d years", y)
			arr = append(arr, y_s)
		}

		if d == 1 {
			d_s = fmt.Sprintf("%d day", d)
			arr = append(arr, d_s)
		} else if d > 1 {
			d_s = fmt.Sprintf("%d days", d)
			arr = append(arr, d_s)
		}

		if h == 1 {
			h_s = fmt.Sprintf("%d hour", h)
			arr = append(arr, h_s)
		} else if h > 1 {
			h_s = fmt.Sprintf("%d hours", h)
			arr = append(arr, h_s)
		}

		if m == 1 {
			m_s = fmt.Sprintf("%d minute", m)
			arr = append(arr, m_s)
		} else if m > 1 {
			m_s = fmt.Sprintf("%d minutes", m)
			arr = append(arr, m_s)
		}

		if s == 1 {
			s_s = fmt.Sprintf("%d second", s)
			arr = append(arr, s_s)
		} else if s > 1 {
			s_s = fmt.Sprintf("%d seconds", s)
			arr = append(arr, s_s)
		}

		if len(arr) == 1 {
			return arr[0]
		}
		if len(arr) == 2 {
			return fmt.Sprintf("%s and %s", arr[0], arr[1])
		}
		if len(arr) == 3 {
			return fmt.Sprintf("%s, %s and %s", arr[0], arr[1], arr[2])
		}
		if len(arr) == 4 {
			return fmt.Sprintf("%s, %s, %s and %s", arr[0], arr[1], arr[2], arr[3])
		}
		return fmt.Sprintf("%s, %s, %s, %s and %s", arr[0], arr[1], arr[2], arr[3], arr[4])
	}

	m, s := seconds/60, seconds%60
	h, m := m/60, m%60
	d, h := h/24, h%24
	y, d := d/365, d%365
	return convert(y, d, h, m, s)
}
