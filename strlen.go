package piscine

func StrLen(s string) int {
	letters_index := []rune(s)
	b := 0
	for c := range letters_index {
		b = c + 1
	}
	return b
}
