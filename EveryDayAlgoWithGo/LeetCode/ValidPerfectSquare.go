package main

func isPerfectSquare(num int) bool {
	for i := 1; i*i <= num; {
		if i*i == num {
			return true
		}
		i++
	}
	return false
}
