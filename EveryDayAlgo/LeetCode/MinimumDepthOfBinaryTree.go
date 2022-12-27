package main

func minDepth(root *TreeNode) int {
	var result int
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	result = minDepth(root.Left)
	if result > minDepth(root.Right) {
		result = minDepth(root.Right)
	}
	return result + 1

}
