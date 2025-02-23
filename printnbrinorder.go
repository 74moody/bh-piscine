package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	arr := []int{}
	if n <= 0 {
		z01.PrintRune('0')
		return
	}
	for n > 0 {
		arr = append(arr, n%10)
		n = n / 10
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	for _, ch := range arr {
		z01.PrintRune(rune(ch + '0'))
	}
}
