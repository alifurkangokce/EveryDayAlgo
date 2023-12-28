package main

func sumOfDigits(n int) int {
	s := 0
	for n > 0 {
		s += n % 10
		n /= 10
	}
	return s
}

func countLargestGroup(n int) int {
	var groups [37]int
	var max, countMax int
	for i := 1; i <= n; i++ {
		s := sumOfDigits(i)
		groups[s]++
		if groups[s] == max {
			countMax++
		} else if groups[s] > max {
			max = groups[s]
			countMax = 1
		}
	}
	return countMax
}
