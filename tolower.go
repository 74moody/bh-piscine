package piscine

func ToLower(s string) string {
	arr := []rune(s)
	for i, a := range arr {
		if a >= 65 && a <= 90 {
			arr[i] = a + 32
		}
	}
	return string(arr)
}
