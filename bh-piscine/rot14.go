package piscine

func Rot14(s string) string {
	d := []rune(s)
	for o, i := range d {
		if i >= 'A' && i < 'M' {
			i = i + 14
			d[o] = i
		} else if i >= 'a' && i < 'm' {
			i = i + 14
			d[o] = i
		} else if i >= 'M' && i <= 'Z' {
			i = i - 12
			d[o] = i
		} else if i >= 'm' && i <= 'z' {
			i = i - 12
			d[o] = i
		}
	}
	return string(d)
}
