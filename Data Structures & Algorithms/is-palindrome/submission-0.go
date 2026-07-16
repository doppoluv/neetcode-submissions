func isLetter(s rune) bool {
	return s >= 'a' && s <= 'z'
}

func isDigit(s rune) bool {
	return s >= '0' && s <= '9'
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	n := 0
	m := len(s) - 1

	for n < len(s) && m >= 0 {
		if !isLetter(rune(s[n])) && !isDigit(rune(s[n])) {
			n++
			continue
		}

		if !isLetter(rune(s[m])) && !isDigit(rune(s[m])) {
			m--
			continue
		}

		if s[n] != s[m] {
			return false
		}

		n++
		m--
	}

	return true
}
