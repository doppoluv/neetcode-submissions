func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	for n := 0; n < len(nums); n++ {
		if n != 0 && nums[n] == nums[n-1] {
			continue
		}

		m := n + 1
		k := len(nums) - 1
		for k > m {
			sum := nums[n] + nums[m] + nums[k]
			if sum == 0 {
				r := [][]int{{nums[n], nums[m], nums[k]}}
				result = append(result, r...)
				m++
				k--

				for m < k && nums[m] == nums[m-1] {
					m++
				}
			}

			if sum > 0 {
				k--
			}

			if sum < 0 {
				m++
			}
		}
	}

	return result
}
