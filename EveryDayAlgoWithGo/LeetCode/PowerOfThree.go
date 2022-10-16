package main

func isPowerOfThree(n int) bool {
	// var x float64
	// var cnt float64
	// if n == 0 {
	// 	return false
	// }
	// for int(x) < n {
	// 	x = math.Pow(3, cnt)
	// 	cnt++
	// }
	// return int(x) == n

	if n < 1 {
		return false
	}

	for n%3 == 0 {
		n /= 3
	}

	return n == 1
}
