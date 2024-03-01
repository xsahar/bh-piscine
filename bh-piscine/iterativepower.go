package piscine

func IterativePower(nb int, power int) int {
	result := 1
	if power < 0 {
		return 0
	}
	for i := power; i > 0; i-- {
		result = result * nb
	}
	return result
}
