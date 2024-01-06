package main

func kLengthApart(nums []int, k int) bool {
	var cnt int
	var check bool
	for _, v := range nums {
		if v == 1 && !check {
			check = true
		} else if v == 1 && check {
			if cnt < k {
				return false
			}
			cnt = 0
		} else {
			cnt++
		}
	}
	return true
}
