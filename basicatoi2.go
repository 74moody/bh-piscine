package piscine

func BasicAtoi2(s string) int {
	arr := []rune(s)
	a := 0
	for _, b := range arr {
		if b >= '0' && b <= '9' {
			a = (a * 10) + int(b) - '0'
		} else {
			return 0
		}
	}
	return a
}
