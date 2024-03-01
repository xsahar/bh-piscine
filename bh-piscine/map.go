package piscine

func Map(f func(int) bool, a []int) []bool {
	var arr []bool
	for _, x := range a {
		r := f(x)
		arr = append(arr, r)
	}
	return arr
}
