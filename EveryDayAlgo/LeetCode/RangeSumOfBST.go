package main

func rangeSumBST(root *TreeNode, low int, high int) int {
	result := 0
	res := []int{}
	rangeSumBSTHelper(root, &res)
	for i := 0; i <= len(res)-1; i++ {
		if res[i] >= low && res[i] <= high {
			result += res[i]
		}
	}
	return result
}
func rangeSumBSTHelper(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	if node.Left != nil {
		rangeSumBSTHelper(node.Left, res)
	}
	if node.Right != nil {
		rangeSumBSTHelper(node.Right, res)
	}
	*res = append(*res, node.Val)
}
