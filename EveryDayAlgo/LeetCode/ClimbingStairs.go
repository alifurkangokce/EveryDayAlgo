package main

func climbStairs(n int) int {
	one, two, temp := 1, 1, 0

	for i := 0; i < n-1; i++ {
		temp = one
		one = one + two
		two = temp
	}
	return one
}
