package main

func sumOddLengthSubarrays(arr []int) int {
	res := 0
	n := len(arr)
	for i, num := range arr {
		cnt := ((i+1)*(n-i) + 1) / 2
		res += num * cnt
	}
	return res
}
