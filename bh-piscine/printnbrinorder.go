package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	digits := [100]int{}
	for n != 0 {
		d := n % 10
		digits[d]++
		n = n / 10
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < digits[i]; j++ {
			z01.PrintRune(rune(i) + '0')
		}
	}
}
