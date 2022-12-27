package main

import "math"

func findTilt(root *TreeNode) int {
	var res int
	valueSum(root, &res)
	return res
}

func valueSum(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	leftSum := valueSum(root.Left, res)
	righSum := valueSum(root.Right, res)
	tilt := math.Abs(float64(leftSum) - float64(righSum))
	*res += int(tilt)
	return root.Val + leftSum + righSum
}
