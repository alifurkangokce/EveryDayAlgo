package main

func findMode(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	max := 0
	res := []int{}
	tmp := make(map[int]int)
	temp := helper2(root, tmp)
	for _, v := range temp {
		if v > max {
			max = v
		}
	}
	for i, v := range temp {
		if v == max {
			res = append(res, i)
		}
	}
	return res
}

func helper2(root *TreeNode, tmp map[int]int) map[int]int {
	if root == nil {
		return tmp
	}
	tmp[root.Val]++
	helper2(root.Left, tmp)
	helper2(root.Right, tmp)
	return tmp
}
