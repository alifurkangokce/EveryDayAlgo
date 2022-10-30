package main

func sumOfLeftLeaves(root *TreeNode) int {
	var result int
	helper3(root, false, &result)
	return result

}
func helper3(root *TreeNode, isLeftNode bool, result *int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil && isLeftNode {
		*result += root.Val
	}
	helper3(root.Left, true, result)
	helper3(root.Right, false, result)
}
