package main

func totalMoney(n int) int {
	var res int
	var sub int
	var base int
	for i := 1; i <= n; i++ {
		if i%7 == 0 {
			base++
			res += sub + 1
			sub = base
		} else {
			sub++
			res += sub
		}

	}
	return res
}
