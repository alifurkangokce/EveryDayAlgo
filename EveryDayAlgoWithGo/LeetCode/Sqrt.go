package main

func mySqrt(x int) int {
	var y int
	for y*y <= x {
		y++
	}
	return y - 1
}
