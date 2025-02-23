package main

import "github.com/01-edu/z01"

type point struct {
	x, y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)

	output := "x = 42, y = 21\n"
	for _, ch := range output {
		z01.PrintRune(ch)
	}
}
