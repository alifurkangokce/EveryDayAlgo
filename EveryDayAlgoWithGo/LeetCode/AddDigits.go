package main

func addDigits(num int) int {
	// for num%10 != num {
	// 	num = (num % 10) + (num-(num%10))/10
	// }
	// return num

	// Best Solution
	if num == 0 {
		return 0
	}
	return 1 + (num-1)%9
}
