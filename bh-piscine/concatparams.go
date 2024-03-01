package piscine

func ConcatParams(args []string) string {
	var s string
	k := 0
	for i := range args {
		i++
		k++
	}
	for i, f := range args {
		s += f
		if i != k-1 {
			s += "\n"
		}
	}
	return s
}
