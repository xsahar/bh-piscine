package piscine

func IsAlpha(s string) bool {
	arr := []rune(s)
	for i := 0; i < len(arr); i++ {
		if (arr[i] < 'a' || arr[i] > 'z') && (arr[i] < 'A' || arr[i] > 'Z') && (arr[i] < '0' || arr[i] > '9') {
			return false
		}
	}
	return true
}
