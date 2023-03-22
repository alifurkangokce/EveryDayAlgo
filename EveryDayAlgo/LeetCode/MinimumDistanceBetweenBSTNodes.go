package main

import "math"

func minDiffInBST(root *TreeNode) int {

	s := []int{}
	helperminDiffInBST(root, &s)
	minDiff := math.MaxInt64

	for i := 1; i < len(s); i++ {
		if s[i]-s[i-1] < minDiff {
			minDiff = s[i] - s[i-1]
		}
	}

	return minDiff

}

func helperminDiffInBST(root *TreeNode, s *[]int) {

	if root == nil {
		return
	}
	helperminDiffInBST(root.Left, s)
	*s = append(*s, root.Val)
	helperminDiffInBST(root.Right, s)

}
