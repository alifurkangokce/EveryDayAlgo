package main

import "math"

func nearestValidPoint(x int, y int, points [][]int) int {
	min := math.MaxFloat64
	var res int = -1
	for k, v := range points {
		valy := math.Abs(float64(y - v[1]))
		valx := math.Abs(float64(x - v[0]))
		if v[0] == x && v[1] == y {
			return k
		} else if v[0] == x {
			if valy < min {
				min = valy
				res = k
			}
		} else if v[1] == y {
			if valx < min {
				min = valx
				res = k
			}
		}
	}
	return res
}
