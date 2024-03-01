package piscine

func SplitWhiteSpaces(s string) []string {
	var st []string
	word := s + " "
	temp := ""
	for _, i := range word {
		if i == ' ' || i == '\n' {
			if temp != "" {
				st = append(st, temp)
				temp = ""
			}
		} else {
			temp = temp + string(i)
		}
	}
	return st
}
