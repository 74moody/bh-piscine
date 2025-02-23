package piscine

func BasicAtoi(s string) int {
	arr := []rune(s)
	a := 0
	for _, b := range arr {
		a = (a * 10) + int(b-'0')
	}
	return a
}
