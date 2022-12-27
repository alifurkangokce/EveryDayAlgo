package main

import "strconv"

func findRelativeRanks(score []int) []string {
	var res []string
	arr, s := BubbleSort(score)

	for _, v := range s {
		cnt := 0
		for i := len(arr) - 1; i >= 0; i-- {
			cnt++
			if arr[i] == v {
				if cnt == 1 {
					res = append(res, "Gold Medal")
					break
				}
				if cnt == 2 {
					res = append(res, "Silver Medal")
					break
				}
				if cnt == 3 {
					res = append(res, "Bronze Medal")
					break
				}
				res = append(res, strconv.Itoa(cnt))
			}
		}

	}

	return res

}
func BubbleSort(array []int) ([]int, []int) {
	var res []int
	res = append(res, array...)
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array, res
}
