package main

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	var res []int
	totalA := 0
	totalB := 0
	setB := make(map[int]bool)
	for i := 0; i <= len(aliceSizes)-1; i++ {
		totalA += aliceSizes[i]
	}
	for i := 0; i <= len(bobSizes)-1; i++ {
		totalB += bobSizes[i]
		setB[bobSizes[i]] = true
	}
	delta := (totalB - totalA) / 2
	for i := 0; i <= len(aliceSizes)-1; i++ {
		if setB[aliceSizes[i]+delta] {
			res = append(res, aliceSizes[i], aliceSizes[i]+delta)
			return res
		}
	}
	return nil
}
