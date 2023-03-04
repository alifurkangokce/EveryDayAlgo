package main

import "strconv"

func selfDividingNumbers(left int, right int) []int {
	var res []int
	for i := left; i <= right; i++ {
		str := strconv.Itoa(i)
		cnt := 0
		for j := 0; j <= len(str)-1; j++ {
			nmbr, _ := strconv.Atoi(string(str[j]))
			if nmbr != 0 && i%nmbr == 0 {
				cnt++
			}
		}
		if cnt == len(str) {
			res = append(res, i)
		}
	}
	return res
}
