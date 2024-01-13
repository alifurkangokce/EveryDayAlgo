package main

func xorOperation(n int, start int) int {
	var number int = start
	for i := 0; i < n-1; i++ {
		start ^= (number + 2)
		number += 2
	}
	return start
}
