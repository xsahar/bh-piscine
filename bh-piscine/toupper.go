package piscine

func ToUpper(s string) string {
	l := []rune(s)
	for i := range l {
		if l[i] >= 'a' && l[i] <= 'z' {
			l[i] -= 32
		}
	}
	return string(l)
}
