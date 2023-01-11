package main

import (
	"strconv"
)

func tree2str(root *TreeNode) string {
	var result string
	helperTree2str(root, &result)
	return result[1 : len(result)-1]
}
func helperTree2str(root *TreeNode, result *string) {
	if root == nil {
		return
	}
	str := strconv.Itoa(root.Val)
	*result += "(" + str

	if root.Left == nil && root.Right != nil {
		*result += "()"
	}
	helperTree2str(root.Left, result)
	helperTree2str(root.Right, result)
	*result += ")"
}
