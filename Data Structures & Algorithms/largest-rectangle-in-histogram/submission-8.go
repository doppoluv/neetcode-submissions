func largestRectangleArea(heights []int) int {
	stack := make([][]int, 0)
	maxS := 0

	for idx, height := range heights {
		h := [][]int{{idx, height}}
		if len(stack) == 0 {
			stack = append(stack, h...)
			maxS = max(maxS, height)
			continue
		}

		lastIdx := idx
		for n := len(stack) - 1; n >= 0; n-- {
			if stack[n][1] > height {
				maxS = max(maxS, stack[n][1]*(idx-stack[n][0]))
				lastIdx = stack[n][0]
				stack = stack[:n]
				continue
			}

			break
		}

		h = [][]int{{lastIdx, height}}
		stack = append(stack, h...)
	}

	// fmt.Println(stack)
	for n := 0; n < len(stack); n++ {
		if len(stack) == 1 {
			maxS = max(maxS, stack[n][1])
			break
		}

		if n == len(stack)-1 {
			maxS = max(maxS, stack[n][1]*(len(heights)-stack[n][0]))
			break
		}

		maxS = max(maxS, stack[n][1]*(stack[len(stack)-1][0]-stack[n][0]+1))
	}

	return maxS
}