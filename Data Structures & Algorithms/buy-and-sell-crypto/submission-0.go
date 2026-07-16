func maxProfit(prices []int) int {
	maxS := 0
	left := 0
	right := 1

	for right < len(prices) {
		if prices[right] >= prices[left] {
			maxS = max(maxS, prices[right]-prices[left])
		} else {
			left = right
		}

		if right == len(prices)-1 {
			break
		}

		right++
	}

	return maxS
}