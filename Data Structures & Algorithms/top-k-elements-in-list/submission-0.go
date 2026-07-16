func topKFrequent(nums []int, k int) []int {
	Map := make(map[int]int)
	for _, num := range nums {
		Map[num]++
	}

	Arr := make(map[int][]int)
	for num, freq := range Map {
		Arr[freq] = append(Arr[freq], num)
	}

	topK := make([]int, k)
	idx := 0
	for n := len(nums) + 1; n > 0; n-- {
		if k <= 0 {
			break
		}

		if Arr[n] != nil {
			for i := 0; i < len(Arr[n]) && k > 0; i++ {
				topK[idx] = Arr[n][i]
				idx++
				k--
			}
		}
	}
	return topK
}