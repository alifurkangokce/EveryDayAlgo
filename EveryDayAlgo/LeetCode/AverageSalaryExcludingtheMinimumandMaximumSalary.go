package main

import "sort"

func average(salary []int) float64 {
	sort.Ints(salary)
	var sum float64
	x := salary[1 : len(salary)-1]
	for _, v := range x {
		sum += float64(v)
	}
	return float64(sum / float64(len(x)))
}
