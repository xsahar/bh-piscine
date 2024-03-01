package piscine

func TrimAtoi(s string) int {
	result := 0
	neg := 1
	for _, v := range s {
		if v >= '0' && v <= '9' {
			result = int(v) - 48 + result*10
		} else if v == '-' && result == 0 {
			neg = -1
		}
	}
	return result * neg
}
