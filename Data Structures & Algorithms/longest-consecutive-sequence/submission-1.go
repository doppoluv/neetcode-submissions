func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	Map := make(map[int]bool)
	for _, num := range nums {
		Map[num] = true
	}

	starts := []int{}
	for num := range Map {
		if Map[num-1] != true {
			starts = append(starts, num)
		}
	}

	longestConsec := 1
	for _, num := range starts {
		consec := 1
		for n := 1; n < len(nums); n++ {
			if Map[num+n] != true {
				break
			}

			consec++
		}

		longestConsec = max(longestConsec, consec)
	}

	return longestConsec
}