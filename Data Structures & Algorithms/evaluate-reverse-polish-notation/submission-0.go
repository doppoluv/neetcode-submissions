func evalRPN(tokens []string) int {
	stack := make([]string, 0)

	for _, s := range tokens {
		if s == "*" || s == "-" || s == "+" || s == "/" {
			if len(stack) < 2 {
				return 0
			}

			num1, _ := strconv.Atoi(stack[len(stack)-2])
			num2, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-2]

			if s == "+" {
				stack = append(stack, fmt.Sprint(num1+num2))
			}

			if s == "-" {
				stack = append(stack, fmt.Sprint(num1-num2))
			}

			if s == "*" {
				stack = append(stack, fmt.Sprint(num1*num2))
			}

			if s == "/" {
				if num2 == 0 {
					return 0
				}
				stack = append(stack, fmt.Sprint(num1/num2))
			}

			continue
		}

		stack = append(stack, s)
	}

	result, _ := strconv.Atoi(stack[0])
	return result
}