package main

func preorderTraversal(root *TreeNode) []int {
	var result []int
	helper(root, &result)
	return result
}

func helper(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)
	helper(root.Left, result)
	helper(root.Right, result)
}
