package main

import "sort"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	var arr []int
	arr = append(arr, root.Val)
	increasingBSTHelper(root, &arr)
	sort.Ints(arr)
	newHead := &TreeNode{Val: arr[0]}
	ptr := newHead
	for i := 1; i <= len(arr)-1; i++ {
		newHead.Right = &TreeNode{Val: arr[i]}
		newHead = newHead.Right
	}
	return ptr
}
func increasingBSTHelper(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	if node.Left != nil {
		*res = append(*res, node.Left.Val)
		increasingBSTHelper(node.Left, res)
	}
	if node.Right != nil {
		*res = append(*res, node.Right.Val)
		increasingBSTHelper(node.Right, res)
	}

}
