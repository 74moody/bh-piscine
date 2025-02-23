package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	output := 1
	for i := nb; i > 0; i-- {
		output = output * i
	}
	return output
}
