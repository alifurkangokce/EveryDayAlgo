package main

import "math"

func intersection2(nums1 []int, nums2 []int) []int {
	map1, map2 := map[int]int{}, map[int]int{}

	for _, ele := range nums1 {
		map1[ele]++
	}
	for _, ele := range nums2 {
		map2[ele]++
	}

	res := []int{}
	for ele := range map1 {
		if map2[ele] > 0 {
			freq := math.Min(float64(map1[ele]), float64(map2[ele]))
			for freq > 0 {
				res = append(res, ele)
				freq--
			}
		}
	}
	return res
}
