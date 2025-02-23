package piscine

func IsLower(s string) bool {
	arr := []rune(s)
	for _, a := range arr {
		if a < 'a' || a > 'z' {
			return false
		}
	}
	return true
}
