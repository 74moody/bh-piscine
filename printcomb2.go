package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for k := i; k <= '9'; k++ {
			z01.PrintRune(i)
			z01.PrintRune(k)

		}
	}
}
