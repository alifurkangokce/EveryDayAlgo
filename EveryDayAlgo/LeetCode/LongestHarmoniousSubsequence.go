package main

func findLHS(nums []int) int {

	res := 0
	tmp := make(map[int]int)

	for _, v := range nums {
		tmp[v]++
	}
	for i, v := range tmp {
		if c, ok := tmp[i+1]; ok {
			if v+c > res {
				res = v + c
			}
		}
	}
	return res
}
