func missingNumber(nums []int) int {
    result := len(nums)

	for i := 0; i < len(nums); i++ {
		result += i - nums[i]
	}

    return result
}