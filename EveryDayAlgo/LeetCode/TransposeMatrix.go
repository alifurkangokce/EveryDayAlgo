package main

func transpose(matrix [][]int) [][]int {
	result := make([][]int, len(matrix[0]))
	for i := 0; i <= len(matrix[0])-1; i++ {
		for j := 0; j <= len(matrix)-1; j++ {
			result[i] = append(result[i], matrix[j][i])
		}
	}
	return result
}
