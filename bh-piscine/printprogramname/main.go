package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	aurg := os.Args[0]
	for i, k := range aurg {
		if i >= 2 {
			z01.PrintRune(k)
		}
	}
	z01.PrintRune('\n')
}
