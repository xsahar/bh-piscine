package piscine

func AlphaCount(s string) int {
	arr := []rune(s)
	count := 0
	v := len(arr)
	for i := 0; i < v; i++ {
		if arr[i] >= 65 && arr[i] <= 90 {
			count++
		}
		if arr[i] >= 97 && arr[i] <= 122 {
			count++
		}
	}
	return count
}
