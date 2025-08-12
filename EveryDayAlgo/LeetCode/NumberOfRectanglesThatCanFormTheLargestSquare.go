package main

func countGoodRectangles(rectangles [][]int) int {
	mm := make(map[int]int)
	for _, v := range rectangles {
		if v[0] < v[1] {
			mm[v[0]]++
		} else {
			mm[v[1]]++
		}
	}
	maxKey := -1
	for k := range mm {
		if k > maxKey {
			maxKey = k
		}
	}
	return mm[maxKey]

}
