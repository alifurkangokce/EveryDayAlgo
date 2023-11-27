package main

func subtractProductAndSum(n int) int {
	var multip, sum int = 1, 0
	for n > 0 {
		n1 := n % 10
		multip *= n1
		sum += n1
		n /= 10
	}
	return multip - sum
}
