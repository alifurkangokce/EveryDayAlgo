package main

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	var res int
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	for i := 0; truckSize != 0 && len(boxTypes) > i; i++ {
		if boxTypes[i][0] < truckSize {
			res += boxTypes[i][0] * boxTypes[i][1]
			truckSize -= boxTypes[i][0]
		} else {
			res += truckSize * boxTypes[i][1]
			return res
		}
	}
	return res
}
