package piscine

func NRune(s string, n int) rune {
	len := len(s)
	a := []rune(s)
	if n > len || n <= 0 {
		return 0
	}
	return a[n-1]
}
