package main

import "fmt"

// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	s1 := nums1[:m]
// 	s2 := nums2[:n]
// 	s1 = append(s1, s2...)
// 	sort.Ints(s1)
// }
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	for k := m + n - 1; k >= 0; k-- {
		if (i >= 0 && j >= 0 && nums1[i] > nums2[j]) || (j < 0) {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
	fmt.Println(nums1)
}
