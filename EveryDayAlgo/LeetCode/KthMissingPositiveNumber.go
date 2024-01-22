package main

func findKthPositive(arr []int, k int) int {
	var cnt int
	var res int
	for i, x := 1, 1; i <= arr[len(arr)-1]; i++ {
		if i != arr[x-1] {
			cnt++
		} else {
			x++
		}
		if cnt == k {
			return i
		}
		res = i
	}
	return (k - cnt) + res
}
