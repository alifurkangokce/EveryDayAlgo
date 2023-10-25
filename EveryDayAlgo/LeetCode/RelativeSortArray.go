package main

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	k := 0
	for _, v := range arr2 {
		for i := k; i < len(arr1); i++ {
			if arr1[i] == v {
				arr1[i], arr1[k] = arr1[k], arr1[i]
				k++
			}
		}
	}
	sort.Ints(arr1[k:])
	return arr1
}
