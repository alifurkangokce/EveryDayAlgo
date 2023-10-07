package main

func canThreePartsEqualSum(arr []int) bool {
	sum := 0

	for _, num := range arr {
		sum += num
	}

	if sum%3 != 0 {
		return false
	}

	partSum, currentSum, count := sum/3, 0, 0

	for i := 0; i < len(arr); i++ {
		currentSum += arr[i]

		if currentSum == partSum {
			currentSum = 0
			count++
		}
	}

	return count >= 3
}
