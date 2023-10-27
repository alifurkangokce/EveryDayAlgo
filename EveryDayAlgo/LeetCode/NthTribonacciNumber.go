package main

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}

	mm := make(map[int]int)
	mm[0] = 0
	mm[1] = 1
	mm[2] = 1
	for i := 3; i <= n; i++ {
		mm[i] = mm[i-3] + mm[i-2] + mm[i-1]
	}
	return mm[n]
}
