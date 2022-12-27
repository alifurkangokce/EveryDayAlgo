package main

func containsNearbyDuplicate(nums []int, k int) bool {
	tmp := make(map[int]int)
	for i, v := range nums {
		if n, ok := tmp[v]; ok && i-n <= k {
			return true
		}
		tmp[v] = i
	}
	return false
}
