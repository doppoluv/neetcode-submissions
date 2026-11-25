func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	lastZero := 0
	for i, n := range nums {
		if n == 0 {
			max = int(math.Max(float64(max), float64(i - lastZero)))
			lastZero = i + 1
		}
	}

	max = int(math.Max(float64(max), float64(len(nums) - lastZero)))
	
	return max
}
