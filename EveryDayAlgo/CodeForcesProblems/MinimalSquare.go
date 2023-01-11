package main

import (
	"fmt"
	"math"
)

func MinimalSquare() {
	var t, a, b int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&a, &b)
		if a <= b {
			fmt.Println(math.Max(math.Pow(float64(a*2), float64(2)), math.Pow(float64(b), float64(2))))
		} else {
			fmt.Println(math.Max(math.Pow(float64(b*2), float64(2)), math.Pow(float64(a), float64(2))))
		}
	}
}
