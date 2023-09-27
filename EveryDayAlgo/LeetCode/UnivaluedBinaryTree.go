package main

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left != nil && root.Left.Val != root.Val {
		return false
	}

	if root.Right != nil && root.Right.Val != root.Val {
		return false
	}

	lVal := isUnivalTree(root.Left)
	if !lVal {
		return false
	}

	rVal := isUnivalTree(root.Right)
	if !rVal {
		return false
	}

	return true
}
