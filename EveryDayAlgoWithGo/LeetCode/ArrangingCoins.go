package main

func arrangeCoins(n int) int {
	x := 0
	i := 1
	res := 0
	for x < n {
		x += i
		i++
		res++
	}
	if x == n {
		return res
	} else {
		return res - 1
	}
}
