func isValidSudoku(board [][]byte) bool {
	var rowSet [9]map[byte]bool
	var colSet [9]map[byte]bool
	var sqrSet [9]map[byte]bool

	// Init maps in sets
	for i := 0; i < 9; i++ {
		rowSet[i] = make(map[byte]bool)
		colSet[i] = make(map[byte]bool)
		sqrSet[i] = make(map[byte]bool)
	}

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			cell := board[row][col]

			if cell == '.' {
				continue
			}

			// mult by 3 to diffirenciate rows
			//
			//(0+0=0) (0+1=1) (0+2=2)
			//(3+0=3) (3+1=4) (3+2=5)
			//(6+0=6) (6+1=7) (6+2=8)
			sqr := (row / 3) * 3 + (col / 3)

			if rowSet[row][cell] || colSet[col][cell] || sqrSet[sqr][cell] {
				return false
			}

			rowSet[row][cell] = true
			colSet[col][cell] = true
			sqrSet[sqr][cell] = true
		}
	}

	return true
}

// Time:  O(n^2)
// Space: O(n^2)