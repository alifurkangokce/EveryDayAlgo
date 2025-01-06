package main

func canFormArray(arr []int, pieces [][]int) bool {
	m := make(map[int][]int)
	for _, piece := range pieces {
		m[piece[0]] = piece
	}
	i := 0
	for i < len(arr) {
		if _, ok := m[arr[i]]; !ok {
			return false
		}
		for _, v := range m[arr[i]] {
			if v != arr[i] {
				return false
			}
			i++
		}
	}
	return true
}
