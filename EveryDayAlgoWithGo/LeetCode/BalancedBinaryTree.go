package main

import "math"

func getDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := getDepth(node.Left)
	right := getDepth(node.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !isBalanced(root.Left) {
		return false
	}

	if !isBalanced(root.Right) {
		return false
	}

	left := getDepth(root.Left)
	right := getDepth(root.Right)
	if math.Abs(float64(left-right)) > 1 {
		return false
	}

	return true
}
