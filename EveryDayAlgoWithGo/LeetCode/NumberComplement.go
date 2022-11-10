package main

func findComplement(num int) int {
	mask := ^0 // mask: ...11111
	for mask&num > 0 {
		mask <<= 1
	}
	return ^num ^ mask
}
