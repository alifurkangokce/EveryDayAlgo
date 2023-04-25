package main

func flipAndInvertImage(image [][]int) [][]int {
	var res [][]int
	var appended []int
	for _, v := range image {
		for i := len(v) - 1; i >= 0; i-- {
			if v[i] == 1 {
				appended = append(appended, 0)
			} else {
				appended = append(appended, 1)
			}
		}
		res = append(res, appended)
		appended = []int{}
	}
	return res
}
