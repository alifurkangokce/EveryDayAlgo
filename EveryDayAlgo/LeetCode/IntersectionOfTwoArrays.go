package main

func intersection(nums1 []int, nums2 []int) []int {
	x := make(map[int]int)
	var res []int
	for i := 0; i < len(nums1); i++ {
		if x[nums1[i]] == 0 {
			x[nums1[i]]++
		}
	}
	for i := 0; i < len(nums2); i++ {
		if x[nums2[i]] == 1 {
			res = append(res, nums2[i])
			x[nums2[i]]++
		}
	}
	return res
}
