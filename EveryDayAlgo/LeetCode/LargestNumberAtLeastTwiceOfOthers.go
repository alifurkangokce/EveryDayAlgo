package main

func dominantIndex(nums []int) int {
	var maxv, idx int
	for i, v := range nums {
		if maxv < v {
			maxv, idx = v, i
		}
	}

	for _, v := range nums {
		if v != maxv && v*2 > maxv {
			return -1
		}
	}

	return idx
}
