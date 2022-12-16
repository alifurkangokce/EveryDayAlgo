package main

import (
	"fmt"
	"sort"
)

func HolidayOfEquality() {
	var inp, req, res int
	var mapReq []int
	fmt.Scan(&inp)
	for i := 0; i < inp; i++ {
		fmt.Scan(&req)
		mapReq = append(mapReq, req)
	}
	sort.Ints(mapReq)
	val := mapReq[len(mapReq)-1]
	for i := 0; i < len(mapReq)-1; i++ {
		res += val - mapReq[i]
	}
	fmt.Print(res)

}
