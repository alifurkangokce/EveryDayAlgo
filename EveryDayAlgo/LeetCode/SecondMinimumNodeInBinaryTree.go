package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findSecondMinimumValue(root *TreeNode) int {
	var queue []*TreeNode
	var node *TreeNode
	var minVal = math.MaxInt64

	queue = append(queue, root)

	for len(queue) != 0 {
		node, queue = queue[0], queue[1:]
		if node.Val != root.Val && node.Val < minVal {
			minVal = node.Val
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
	}

	if minVal != math.MaxInt64 {
		return minVal
	} else {
		return -1
	}
}
