package main

func tictactoe(moves [][]int) string {

	const Winner = 3

	col := [2][3]int{{0, 0, 0}, {0, 0, 0}}
	row := [2][3]int{{0, 0, 0}, {0, 0, 0}}
	diag := [2][2]int{{0, 0}, {0, 0}}

	for index, value := range moves {
		var player = index % 2
		col[player][value[0]]++
		row[player][value[1]]++

		if value[0] == value[1] {
			diag[player][0]++
		}

		if value[0]+value[1] == 2 {
			diag[player][1]++
		}

		if col[player][0] == Winner || col[player][1] == Winner || col[player][2] == Winner || row[player][0] == Winner || row[player][1] == Winner || row[player][2] == Winner || diag[player][0] == Winner || diag[player][1] == Winner {
			if player == 0 {
				return "A"
			}
			return "B"
		}
	}

	if len(moves) == 9 {
		return "Draw"
	}

	return "Pending"
}
