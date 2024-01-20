package main

func countGoodTriplets(arr []int, a int, b int, c int) int {
	var res int
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			for k := j + 1; k <= len(arr)-1; k++ {
				if absCountGoodTriplets(arr[i]-arr[j]) <= a && absCountGoodTriplets(arr[j]-arr[k]) <= b && absCountGoodTriplets(arr[i]-arr[k]) <= c {
					res++
				}
			}
		}
	}
	return res
}
func absCountGoodTriplets(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
