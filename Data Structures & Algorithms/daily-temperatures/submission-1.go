func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)

	for idx, temp := range temperatures {
		if len(stack) == 0 {
			stack = append(stack, idx)
			continue
		}

		count := 0
		for n := len(stack) - 1; n >= 0; n-- {
			if temp > temperatures[stack[n]] {
				temperatures[stack[n]] = idx - stack[n]
				count++
				continue
			}

            break
		}

		stack = stack[:len(stack)-count]
		stack = append(stack, idx)
	}

	if len(stack) != 0 {
		for _, idx := range stack {
			temperatures[idx] = 0
		}
	}

	return temperatures
}