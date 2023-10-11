package main

var res int

func sumRootToLeaf(root *TreeNode) int {
	res = 0
	sumRootToLeafHelper(root, 0)
	return res
}

func sumRootToLeafHelper(root *TreeNode, n int) {
	if root.Left == nil && root.Right == nil {
		res += n*2 + root.Val
		return
	}
	if root.Left != nil {
		sumRootToLeafHelper(root.Left, n*2+root.Val)
	}
	if root.Right != nil {
		sumRootToLeafHelper(root.Right, n*2+root.Val)
	}
}
