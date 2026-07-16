func lengthOfLongestSubstring(s string) int {
	maxS := 0
	left := 0
	right := 0
	arr := make(map[byte]int)

	for right < len(s) {
		if arr[s[right]] == 0 {
			arr[s[right]]++
		} else {
			maxS = max(maxS, right-left)

			for s[left] != s[right] {
				arr[s[left]]--
				left++
			}

			if left != right {
				left++
			}
		}

		if right == len(s)-1 {
			maxS = max(maxS, right-left+1)
		}

		right++
	}

	return maxS
}