package main

func minOperations(logs []string) int {
	move := 0
	for _, v := range logs {
		if v == "../" {
			if move != 0 {
				move--
			}
		} else if v != "./" {
			move++
		}
	}
	return move
}
