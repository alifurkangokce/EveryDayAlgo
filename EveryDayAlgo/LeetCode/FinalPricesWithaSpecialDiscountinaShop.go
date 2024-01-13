package main

func finalPrices(prices []int) []int {
	var res []int
	for i := 0; i <= len(prices)-1; i++ {
		var check bool
		for j := i + 1; j <= len(prices)-1; j++ {
			if prices[i] >= prices[j] {
				res = append(res, prices[i]-prices[j])
				check = true
				break
			}
		}
		if !check {
			res = append(res, prices[i])
		}
	}
	return res
}
