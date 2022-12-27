package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepthn(root *Node) int {
	max := 0
	dfsn(root, 1, &max)
	return max

}

func dfsn(root *Node, depth int, max_depth *int) {
	if root == nil {
		return
	}

	*max_depth = maxn(depth, *max_depth)
	for _, child := range root.Children {
		dfsn(child, depth+1, max_depth)
	}
}

func maxn(x, y int) int {
	if x < y {
		return y
	}
	return x
}
