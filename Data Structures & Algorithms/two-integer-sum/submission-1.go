func twoSum(nums []int, target int) []int {
	Map := make(map[int]int)
	for idx, num := range nums {
		reverseNum := target - num
		if Map[reverseNum] != 0 {
			return []int{Map[reverseNum] - 1, idx}
		}

		Map[num] = idx + 1
	}

	return []int{}
}