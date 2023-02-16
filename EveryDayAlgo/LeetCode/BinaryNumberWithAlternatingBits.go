package main

func hasAlternatingBits(n int) bool {
	lastR := -1

	for n > 0 {
		r := n % 2
		n = n / 2

		if r == lastR {
			return false
		}

		lastR = r
	}

	return true
}
