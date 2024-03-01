package piscine

func Abort(a, b, c, d, e int) int {
	tab := []int{a, b, c, d, e}
	for i := 0; i < len(tab)-1; i++ {
		for j := i + 1; j < len(tab); j++ {
			if tab[i] > tab[j] {
				tab[i], tab[j] = tab[j], tab[i]
			}
		}
	}
	return tab[2]
}
