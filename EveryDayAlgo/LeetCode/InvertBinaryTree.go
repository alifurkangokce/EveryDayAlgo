package main

func invertTree(root *TreeNode) *TreeNode {
	dfs(root)
	return root
}
func dfs(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil {
		dfs(root.Left)
	}

	if root.Right != nil {
		dfs(root.Right)
	}

	root.Left, root.Right = root.Right, root.Left
}
