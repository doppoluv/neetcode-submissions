func isValidSudoku(board [][]byte) bool {
	MapStr := make(map[byte][]bool)
	MapCol := make(map[byte][]bool)
	MapSquare := make(map[byte][]bool)
	for n := 0; n < len(board); n++ {
		for m := 0; m < len(board); m++ {
			b := board[n][m]
			if b == '.' {
				continue
			}

			if MapStr[b] == nil {
				MapStr[b] = make([]bool, len(board))
			}
			if MapStr[b][n] {
				return false
			}
			MapStr[b][n] = true

			if MapCol[b] == nil {
				MapCol[b] = make([]bool, len(board))
			}
			if MapCol[b][m] {
				return false
			}
			MapCol[b][m] = true

			if MapSquare[b] == nil {
				MapSquare[b] = make([]bool, len(board))
			}
			if MapSquare[b][(n/3)*3+(m/3)] {
				return false
			}
			MapSquare[b][(n/3)*3+(m/3)] = true
		}
	}

	return true
}
