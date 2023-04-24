package main

func largeGroupPositions(s string) [][]int {
	var sameCounter int
	var arrRange [][]int
	startindex := -1

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			sameCounter++
			if startindex < 0 {
				startindex = i
			}
		} else {
			if sameCounter >= 2 {
				arrRange = append(arrRange, []int{startindex, i})
			}
			sameCounter = 0
			startindex = -1
		}
	}
	if sameCounter >= 2 {
		arrRange = append(arrRange, []int{startindex, len(s) - 1})
	}
	return arrRange

}
