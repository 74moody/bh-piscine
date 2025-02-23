package piscine

func StrRev(s string) string {
	arr := []rune(s)
	a := ""
	for index := len(arr) - 1; index >= 0; index-- {
		a += string(arr[index])
	}
	return a
}
