package main

func maxNumberOfBalloons(text string) int {
	var res int
	var control bool
	var balloon string = "balloon"
	mB := make(map[rune]int)
	for _, v := range balloon {
		mB[v]++
	}
	mI := make(map[rune]int)
	for _, v := range text {
		mI[v]++
	}
	for !control {
		for k, v := range mB {
			if mI[k] < v {
				control = false
				return res

			}
			mI[k] -= v
		}
		res++
	}
	return res
}
