package main

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	arr1 := make([]int, 0)
	arr2 := make([]int, 0)

	addToArray(root1, &arr1)
	addToArray(root2, &arr2)

	if len(arr1) != len(arr2) {
		return false
	}

	for index, _ := range arr2 {
		if arr1[index] != arr2[index] {
			return false
		}
	}

	return true
}

func addToArray(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*arr = append(*arr, root.Val)
	}
	addToArray(root.Left, arr)
	addToArray(root.Right, arr)
}
