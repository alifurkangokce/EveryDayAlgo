package main

import (
	"fmt"
	"math"
)

func BeautifulMatrix() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			var inp int
			fmt.Scan(&inp)
			if inp == 1 {
				fmt.Println(math.Abs(float64(2-i)) + math.Abs(float64(2-j)))
			}
		}
	}

}
