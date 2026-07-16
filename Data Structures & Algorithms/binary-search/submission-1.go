func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	idx := (right - left + 1) / 2

	for (right-left+1) > 1 && left <= right {
		if nums[idx] == target {
			return idx
		}

		if nums[idx] < target {
			left = idx + 1
		} else {
			right = idx - 1
		}
		idx = left + (right-left+1)/2

		fmt.Println(left, right, idx)
	}

	if idx < len(nums) && idx >= 0 && nums[idx] == target {
		return idx
	}

	return -1
}
