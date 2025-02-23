package piscine

var convert string

func Atoi(s string) int {
	a := 0
	arr := []rune(s)
	negative := false
	start := 0
	if len(arr) == 0 {
		return 0
	}
	if arr[0] == '-' {
		negative = true
		start++
	} else if arr[0] == '+' {
		start++
	}
	for i, b := range arr {
		if i == 0 {
			continue
		}
		if b >= '0' && b <= '9' {
			a = (a * 10) + int(b) - '0'
		} else {
			return 0
		}
	}
	if negative {
		return -a
	}
	return a
}
