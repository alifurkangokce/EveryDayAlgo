package main

func postorder(root *Node) []int {
	var result []int
	if root == nil {
		return result
	}

	postOrderHelper(root, &result)
	result = append(result, root.Val)
	return result
}
func postOrderHelper(root *Node, result *[]int) {
	if root == nil {
		return
	}
	for _, v := range root.Children {
		postOrderHelper(v, result)
		*result = append(*result, v.Val)
	}
}
