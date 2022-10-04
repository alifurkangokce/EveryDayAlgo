package main

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	res := []string{}
	tracePath(root, "", &res)
	return res
}

func tracePath(node *TreeNode, current string, res *[]string) {
	if current == "" {
		current = strconv.Itoa(node.Val)
	} else {
		current = current + "->" + strconv.Itoa(node.Val)
	}
	if node.Left == nil && node.Right == nil {
		*res = append(*res, current)
	}
	if node.Left != nil {
		tracePath(node.Left, current, res)
	}
	if node.Right != nil {
		tracePath(node.Right, current, res)
	}
}
