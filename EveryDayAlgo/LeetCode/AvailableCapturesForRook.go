package main

func numRookCaptures(board [][]byte) int {
	var r1, r2 int
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 'R' {
				r1 = i
				r2 = j
				break
			}
		}
	}
	var result int

	isp := false
	for i := 0; i <= r1; i++ {
		if board[r1][i] == 'p' {
			isp = true

		}
		if board[r1][i] == 'B' {

			isp = false
		}
	}
	if isp {
		result++
	}
	isp = false

	for i := 7; i >= r1; i-- {
		if board[r1][i] == 'p' {
			isp = true

		}
		if board[r1][i] == 'B' {

			isp = false
		}
	}
	if isp {
		result++
	}

	isp = false
	for i := 0; i <= r2; i++ {
		if board[i][r2] == 'p' {
			isp = true

		}
		if board[i][r2] == 'B' {

			isp = false
		}
	}
	if isp {
		result++
	}
	isp = false

	for i := 7; i >= r2; i-- {
		if board[i][r2] == 'p' {
			isp = true

		}
		if board[i][r2] == 'B' {

			isp = false
		}
	}
	if isp {
		result++
	}

	return result
}
