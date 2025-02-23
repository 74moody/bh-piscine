package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	letters := []rune(s)
	for _, b := range letters {
		z01.PrintRune(b)
	}
}
