package piscine

func NRune(s string, n int) rune {
	arr := []rune(s)
	if n > len(s) || n < 0 || n-1 < 0 {
		return 0
	} else {
		return arr[n-1]
	}
}
