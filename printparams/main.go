package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.args[1:]
	for _, argz := range args {
		z01.PrintRune(args)
	}
}
