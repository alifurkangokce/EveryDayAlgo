package main

import "sort"

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var result []int
	min := 0
	helperPost(root, &result)
	if len(result) == 1 {
		return result[0]
	} else {
		sort.Ints(result)
		min = result[1] - result[0]
		for i := 2; i <= len(result)-1; i++ {
			if result[i]-result[i-1] < min {
				min = result[i] - result[i-1]
			}
		}
	}
	return min

}
