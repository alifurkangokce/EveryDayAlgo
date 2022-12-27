package main

func isHappy(n int) bool {
	m := map[int]bool{}
	for n != 1 && !m[n] {
		m[n] = true
		var t int
		for ; n != 0; n /= 10 {
			d := n % 10
			t += d * d
		}
		n = t
	}
	return n == 1
}
