package main

import "math"

func minTimeToVisitAllPoints(points [][]int) int {
	var res float64
	for i := 0; i < len(points)-1; i++ {
		absdif1 := math.Abs(float64(points[i][0]) - float64(points[i+1][0]))
		absdif2 := math.Abs(float64(points[i][1]) - float64(points[i+1][1]))
		if absdif1 >= absdif2 {
			res += absdif1
		} else {
			res += absdif2
		}
	}
	return int(res)
}
