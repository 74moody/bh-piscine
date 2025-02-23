package piscine

func AlphaCount(s string) int {
	count := 0
	arr := []rune(s)
	for _, v := range arr {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
			count++
		}
	}
	return count
}
