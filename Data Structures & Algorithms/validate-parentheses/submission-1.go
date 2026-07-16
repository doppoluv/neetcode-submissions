func isValid(s string) bool {
	stack := make([]rune, len(s))
	openParen := 0
	for _, c := range s {
		if c == '}' {
			if len(stack) < 1 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
			openParen--
			continue
		}
		if c == ']' {
			if len(stack) < 1 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
			openParen--
			continue
		}
		if c == ')' {
			if len(stack) < 1 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
			openParen--
			continue
		}

		stack = append(stack, c)
		openParen++
	}

	return openParen == 0
}
