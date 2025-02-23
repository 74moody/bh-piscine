package piscine

func IsPrintable(s string) bool {
	arr := []rune(s)
	for _, a := range arr {
		if a < 32 || a > 126 {
			return false
		}
	}
	return true
}
