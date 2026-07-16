func maxArea(heights []int) int {
	n := 0
	m := len(heights) - 1

	maxS := (m - n) * min(heights[m], heights[n])
	for n < m {
		S := (m - n) * min(heights[m], heights[n])
		maxS = max(maxS, S)

		if heights[n] > heights[m] {
			m--
			continue
		}

		if heights[n] <= heights[m] {
			n++
			continue
		}
	}

	return maxS
}