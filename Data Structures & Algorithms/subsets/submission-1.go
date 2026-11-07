func subsets(nums []int) [][]int {
	subsets := [][]int{{}}

	for i, num := range nums {
		step := int(math.Pow(2.0, float64(i)))
		for j := 0; j < step; j++ {
			subset := append([]int{}, subsets[j]...)
			subset = append(subset, num)
			subsets = append(subsets, subset)
		}
	}

	return subsets
}
