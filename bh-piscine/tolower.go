package piscine

func ToLower(s string) string {
	l := []rune(s)
	for i := range l {
		if l[i] >= 'A' && l[i] <= 'Z' {
			l[i] += 32
		}
	}
	return string(l)
}
