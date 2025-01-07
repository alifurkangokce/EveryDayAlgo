package main

func getMaximumGenerated(n int) int {
	m := make(map[int]int)
	m[0] = 0
	m[1] = 1
	if n == 0 || n == 1 {
		return m[n]
	}
	max := 0
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			m[i] = m[i/2]
		} else {
			m[i] = m[i/2] + m[i/2+1]
		}
		if m[i] > max {
			max = m[i]
		}
	}
	return max
}
