package main

func distributeCandies(candyType []int) int {
	only := (len(candyType) / 2)
	x := make(map[int]int)
	for _, v := range candyType {
		x[v]++
	}
	if only <= len(x) {
		return only
	} else {
		return len(x)
	}
}
