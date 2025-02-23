package piscine

func IsAlpha(s string) bool {
	arr := []rune(s)
	for _, a := range arr {
		if (a < 'a' || a > 'z') && (a < 'A' || a > 'Z') && (a < '0' || a > '9') {
			return false
		}
	}
	return true
}
