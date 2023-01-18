package main

func findTarget(root *TreeNode, k int) bool {
	var nums []int
	findHelper(root, &nums)
	m := make(map[int]int)
	for index, number := range nums {
		if _, ok := m[k-number]; ok {
			return true
		}
		m[number] = index
	}
	return false
}
func findHelper(root *TreeNode, m *[]int) {
	if root == nil {
		return
	}
	*m = append(*m, root.Val)
	findHelper(root.Left, m)
	findHelper(root.Right, m)
}
