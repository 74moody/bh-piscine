package piscine

func IsNumeric(s string) bool {
	arr := []rune(s)
	for _, a := range arr {
		if a < '0' || a > '9' {
			return false
		}
	}
	return true
}
