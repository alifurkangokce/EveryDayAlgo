package main

import "strconv"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	var ranges []string
	var start, end int

	for i := 0; i < len(nums); i++ {
		if i-start == nums[i]-nums[start] {
			end = i
		} else {
			ranges = append(ranges, generateRange(nums[start], nums[end]))
			start = i
			end = i
		}
	}
	ranges = append(ranges, generateRange(nums[start], nums[end]))

	return ranges
}

func generateRange(start, end int) string {
	if start == end {
		return strconv.Itoa(start)
	}
	return strconv.Itoa(start) + "->" + strconv.Itoa(end)
}
