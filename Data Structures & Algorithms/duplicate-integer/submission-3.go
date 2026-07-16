func hasDuplicate(nums []int) bool {
	Map := make(map[int]struct{})
	for _, num := range nums {
		Map[num] = struct{}{}
	}

	if len(Map) != len(nums) {
		return true
	}

	return false
}