func pivotIndex(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	left := 0
	for i, n := range nums {
		temp := sum - n
		if temp % 2 == 0 && left == temp / 2 {
			return i
		}

		left += n
	}

	return -1
}
