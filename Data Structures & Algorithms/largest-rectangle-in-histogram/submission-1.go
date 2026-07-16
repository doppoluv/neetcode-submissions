func largestRectangleArea(heights []int) int {
	stack := make([]int, 0)
	maxS := 0
	maxCurrent := 0

	for _, height := range heights {
		if height == 0 {
			stack = stack[:0]
			continue
		}

		if len(stack) == 0 {
			stack = append(stack, height)
			maxCurrent = height
            maxS = max(maxS, maxCurrent)
			continue
		}

		stack = append(stack, height)
		for n := len(stack) - 1; n >= 0; n-- {
			if stack[n] < height {
				height = stack[n]
			}
			maxCurrent = max(maxCurrent, (len(stack)-n)*height)
		}

		maxS = max(maxS, maxCurrent)
	}

	return maxS
}
