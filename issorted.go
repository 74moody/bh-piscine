package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	n := len(a)
	if n <= 1 {
		return true
	}
	ascendingSorted := true
	for i := 0; i < n-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			ascendingSorted = false
			break
		}
	}
	descendingSorted := true
	for i := 0; i < n-1; i++ {
		if f(a[i], a[i+1]) < 0 {
			descendingSorted = false
			break
		}
	}
	return ascendingSorted || descendingSorted
}
