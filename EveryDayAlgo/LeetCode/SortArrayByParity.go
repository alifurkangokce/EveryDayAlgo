package main

func sortArrayByParity(nums []int) []int {
	var arr1 []int
	var arr2 []int
	for _, v := range nums {
		if v%2 == 0 {
			arr1 = append(arr1, v)
		} else {
			arr2 = append(arr2, v)
		}
	}
	return append(arr1, arr2...)
}
