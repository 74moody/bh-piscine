package piscine

func ToUpper(s string) string {
	arr := []rune(s)
	for i, a := range arr {
		if a >= 97 && a <= 122 {
			arr[i] = a + 32
		}
	}
	return string(arr)
}
