package main

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	min := prices[0]
	result := prices[0] - prices[1]
	for i := 0; i <= len(prices)-1; i++ {
		if min > prices[i] {
			min = prices[i]
		}
		if min-prices[i] < result {
			result = min - prices[i]
		}
	}
	if result < 0 {
		return result * -1
	}
	return 0
}
