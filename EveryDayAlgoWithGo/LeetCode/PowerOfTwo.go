package main

import "fmt"

func isPowerOfTwo(n int) bool {
	// var x float64
	// var cnt float64
	// if n == 0 {
	// 	return false
	// }
	// for int(x) < n {
	// 	x = math.Pow(2, cnt)
	// 	cnt++
	// }
	// return int(x) == n

	x := n & (n - 1)
	fmt.Print(x)
	return n != 0 && n&(n-1) == 0
}
