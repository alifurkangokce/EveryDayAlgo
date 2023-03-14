package main

func countPrimeSetBits(left int, right int) int {
	res := 0
	for i := left; i <= right; i++ {
		n := i
		cnt := 0
		for n > 0 {
			cnt += n & 1
			n = n >> 1
		}
		if countPrimeHelper(cnt) {
			res++
		}
	}
	return res
}
func countPrimeHelper(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}
