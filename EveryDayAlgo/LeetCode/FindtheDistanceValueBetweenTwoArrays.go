package main

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	var res int
	for _, v := range arr1 {
		check := true
		for _, v2 := range arr2 {
			if abs(v-v2) <= d {
				check = false
				break
			}
		}
		if check {
			res++
		}
	}
	return res
}
