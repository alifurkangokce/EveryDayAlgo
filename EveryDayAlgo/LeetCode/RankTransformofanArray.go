package main

import "sort"

func arrayRankTransform(arr []int) []int {

	arr2 := make([]int, len(arr))
	copy(arr2, arr)
	mm := make(map[int]int)
	sort.Ints(arr2)
	i := 1
	for _, v := range arr2 {
		if mm[v] == 0 {
			mm[v] = i
			i++
		}

	}
	var res []int
	for i := 0; i <= len(arr)-1; i++ {
		res = append(res, mm[arr[i]])
	}
	return res
}
