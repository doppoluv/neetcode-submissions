func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	result[0] = 1

	for n := 1; n < len(nums); n++ {
		result[n] = result[n-1] * nums[n-1]
	}

	for n := len(nums) - 2; n >= 0; n-- {
		result[n] = result[n] * nums[n+1]
		nums[n] = nums[n] * nums[n+1]
	}

	return result
}
