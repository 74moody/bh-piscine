package piscine

func Index(s string, toFind string) int {
	sl := len(s)
	tl := len(toFind)
	for i := 0; i < sl-tl; i++ {
		if s[i:i+tl] == toFind {
			return i
		}
	}
	return -1
}
