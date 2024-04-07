package main

import "sort"

func frequencySort(nums []int) []int {
	var res []int

	freMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		freMap[nums[i]]++
	}

	var arr [][2]int

	for num, fre := range freMap {
		arr = append(arr, [2]int{num, fre})
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i][1] == arr[j][1] {
			return arr[i][0] > arr[j][0]
		}

		return arr[i][1] < arr[j][1]
	})

	for i := 0; i < len(arr); i++ {
		for j := 0; j < arr[i][1]; j++ {
			res = append(res, arr[i][0])
		}
	}

	return res
}
