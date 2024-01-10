package main

func canBeEqual(target []int, arr []int) bool {
	freq := [1001]int{}
	for _, v := range arr {
		freq[v]++
	}
	for _, v := range target {
		freq[v]--
	}
	for _, v := range freq {
		if v != 0 {
			return false
		}
	}
	return true
}
