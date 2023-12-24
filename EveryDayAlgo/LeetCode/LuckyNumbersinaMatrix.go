package main

func luckyNumbers(matrix [][]int) []int {
	answerArray := []int{}
	indexOfMin := 0
	for row := 0; row < len(matrix); row++ {
		min := 100000
		for col := 0; col < len(matrix[row]); col++ {
			if matrix[row][col] < min {
				min = matrix[row][col]
				indexOfMin = col
			}
		}
		max := -1
		for col := indexOfMin; col >= 0; col++ {
			for row := 0; row < len(matrix); row++ {
				if matrix[row][col] > max {
					max = matrix[row][col]
				}
			}
			break
		}
		if max == min {
			answerArray = append(answerArray, min)
		}
	}
	return answerArray
}
