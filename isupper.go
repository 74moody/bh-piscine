package piscine

func IsUpper(s string) bool {
	arr := []rune(s)
	for _, a := range arr {
		if a < 'A' || a > 'Z' {
			return false
		}
	}
	return true
}
