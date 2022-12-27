package main

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if (nil == root) != (nil == subRoot) {
		return false
	}
	if isSame(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSame(root *TreeNode, subRoot *TreeNode) bool {
	if nil == root && nil == subRoot {
		return true
	}
	if nil == root || nil == subRoot {
		return false
	}
	return root.Val == subRoot.Val && isSame(root.Left, subRoot.Left) && isSame(root.Right, subRoot.Right)
}
