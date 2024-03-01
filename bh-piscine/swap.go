package piscine

func Swap(a *int, b *int) {
	r := *a
	j := *b
	*a = j
	*b = r
}
