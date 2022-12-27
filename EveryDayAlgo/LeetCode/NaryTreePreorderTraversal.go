package main

func preorder(root *Node) []int {
	var result []int
	if root == nil {
		return result
	}

	result = append(result, root.Val)
	childHelper(root, &result)
	return result
}

func childHelper(root *Node, result *[]int) {
	if root == nil {
		return
	}
	for _, v := range root.Children {
		*result = append(*result, v.Val)
		childHelper(v, result)
	}
}
