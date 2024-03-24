package main

import "sort"

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	numbersToBeRemoved := int(float64(len(arr)) * 0.1)
	arr = arr[numbersToBeRemoved/2 : len(arr)-numbersToBeRemoved/2]
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return float64(sum) / float64(len(arr))

}
