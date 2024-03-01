package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printInt(x rune) {
	z01.PrintRune(x)
}

func main() {
	points := &point{}

	setPoint(points)

	for _, ch := range "x = " {
		z01.PrintRune(ch)
	}
	printInt(52)
	printInt(50)
	for _, ch := range ", y = " {
		z01.PrintRune(ch)
	}
	printInt(50)
	printInt(49)
	z01.PrintRune('\n')
}
