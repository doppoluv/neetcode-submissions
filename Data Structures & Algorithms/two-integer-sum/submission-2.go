func twoSum(nums []int, target int) []int {
	Map := make(map[int]int)
	for idx, num := range nums {
		reverseNum := target - num
		if _, found := Map[reverseNum]; found {
			return []int{Map[reverseNum], idx}
		}

		Map[num] = idx
	}

	return []int{}
}