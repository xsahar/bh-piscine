package piscine

func LastRune(s string) rune {
	arr := []rune(s)
	lenght := len(arr)
	return arr[lenght-1]
}
