package main

func postorderTraversal(root *TreeNode) []int {
	var result []int
	helperPost(root, &result)
	return result
}

func helperPost(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	helperPost(root.Left, result)

	helperPost(root.Right, result)
	*result = append(*result, root.Val)

}
