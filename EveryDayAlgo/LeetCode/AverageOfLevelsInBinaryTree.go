package main

func averageOfLevels(root *TreeNode) []float64 {
	result := []float64{}
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		newQueue := []*TreeNode{}
		level := len(result)
		result = append(result, 0)
		for _, node := range queue {
			result[level] += float64(node.Val)
			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}
			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
		}
		result[level] /= float64(len(queue))
		queue = newQueue
	}
	return result
}
