package piscine

func Capitalize(s string) string {
	first_letter := true
	word := ""
	for _, ch := range word {
		if ch >= '0' && ch <= '9' || ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
			if first_letter {
				if ch >= 'a' && ch <= 'z' {
					ch = ch - 32
				}
				first_letter = false
			} else {
				if ch >= 'A' && ch <= 'Z' {
					ch = ch + 32
				}
			}
		} else {
			first_letter = true
		}
		word = word + string(ch)
	}
	return word
}
