package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	for i := len(arg) - 1; i >= 0; i-- {
		word := arg[i]
		for _, c := range word {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}
