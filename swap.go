package piscine

func Swap(a *int, b *int) {
	m := *a
	n := *b
	*a = n
	*b = m
}
