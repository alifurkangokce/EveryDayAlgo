package main

func countBalls(lowLimit int, highLimit int) int {
	mm := make(map[int]int)
	for i := lowLimit; i <= highLimit; i++ {
		total := 0
		ll := i
		for ll > 0 {
			total += ll % 10
			ll /= 10
		}
		mm[total]++
	}
	max := 0
	for _, v := range mm {
		if v > max {
			max = v
		}
	}
	return max
}
