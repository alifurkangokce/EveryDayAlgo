package main

import "fmt"

func GeorgeAndAccommodation() {
	var inp, cnt int
	fmt.Scan(&inp)
	for i := 0; i < inp; i++ {
		var p, q int
		fmt.Scan(&p, &q)
		if q-p >= 2 {
			cnt++
		}
	}
	fmt.Print(cnt)
}
