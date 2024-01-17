package main

func numWaterBottles(numBottles int, numExchange int) int {
	var res int = numBottles

	for numBottles >= numExchange {
		diff := numBottles / numExchange
		res += diff
		modd := numBottles % numExchange
		numBottles = diff + modd
	}
	return res
}
