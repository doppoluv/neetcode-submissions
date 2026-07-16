func searchMatrix(matrix [][]int, target int) bool {
	left := 0
	right := len(matrix) - 1
	idx := (right - left + 1) / 2

	for left <= right {
		if (idx+1 == len(matrix) && matrix[idx][0] <= target) || (idx+1 < len(matrix) && matrix[idx][0] <= target && matrix[idx+1][0] > target) {
			left = 0
			right = len(matrix[idx]) - 1
			idxN := (right - left + 1) / 2

			for left <= right {
				if matrix[idx][idxN] == target {
					return true
				}

				if matrix[idx][idxN] > target {
					right = idxN - 1
				}

				if matrix[idx][idxN] < target {
					left = idxN + 1
				}

				idxN = left + (right-left+1)/2
			}

			if idxN >= 0 && idxN < len(matrix[idx]) && matrix[idx][idxN] == target {
				return true
			}

			return false
		}

		if matrix[idx][0] > target {
			right = idx - 1
		}

		if matrix[idx][0] < target {
			left = idx + 1
		}

		idx = left + (right-left+1)/2
	}

	return false
}
