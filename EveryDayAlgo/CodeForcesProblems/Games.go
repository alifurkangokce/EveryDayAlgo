package main

import "fmt"

func Games() {
	var in, x, y, res int
	fmt.Scan(&in)
	var xA, yA []int
	for i := 0; i < in; i++ {
		fmt.Scan(&x, &y)
		xA = append(xA, x)
		yA = append(yA, y)
	}
	for i := 0; i < len(xA); i++ {
		for j := 0; j < len(yA); j++ {
			if xA[i] == yA[j] {
				res++
			}
		}
	}
	fmt.Print(res)
}
