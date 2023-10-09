package main

func prefixesDivBy5(nums []int) (res []bool) {
	r := 0
	for _, v := range nums {
		r = (r*2 + v) % 5
		res = append(res, r == 0)
	}
	return res
}
