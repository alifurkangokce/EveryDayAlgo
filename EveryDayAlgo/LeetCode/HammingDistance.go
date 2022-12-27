package main

func hammingDistance(x int, y int) int {
	var ans int
	z := x ^ y
	for z > 0 {
		ans += z & 1 // Every loop find last one
		z >>= 1
	}
	return ans
}
