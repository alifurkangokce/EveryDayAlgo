package main

import "golang.org/x/exp/slices"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var res []bool
	m := slices.Max(candies)
	for _, v := range candies {
		if v+extraCandies >= m {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	return res
}
