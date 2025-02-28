package main

func numberOfMatches(n int) int {
	var res int
	for n > 1 {
		mod := n % 2
		if mod == 1 {
			res += (n - 1) / 2
			n = (n-1)/2 + 1
		} else {
			res += n / 2
			n = n / 2
		}
	}
	return res
}
