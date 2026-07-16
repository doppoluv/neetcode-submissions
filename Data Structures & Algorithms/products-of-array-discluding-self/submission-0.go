func productExceptSelf(nums []int) []int {
	simple := make([]int, len(nums))
	simple[0] = 1
	simple[1] = nums[0]
	for n := 2; n < len(nums); n++ {
		simple[n] = simple[n-1] * nums[n-1]
	}

	reverse := make([]int, len(nums))
	reverse[len(nums)-1] = 1
	reverse[len(nums)-2] = nums[len(nums)-1]
	for n := len(nums) - 3; n >= 0; n-- {
		reverse[n] = reverse[n+1] * nums[n+1]
	}

	result := make([]int, len(nums))
	for n := 0; n < len(nums); n++ {
		result[n] = simple[n] * reverse[n]
	}

	return result
}