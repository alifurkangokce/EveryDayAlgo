package main

import (
	"fmt"
	"strings"
)

func FoxAndSnake() {
	var in1, in2 int
	fmt.Scan(&in1, &in2)
	for i := 1; i <= in1; i++ {
		if i%4 == 0 {
			fmt.Println("#" + strings.Repeat(".", in2-1))
		} else if i%2 == 0 {
			fmt.Println(strings.Repeat(".", in2-1) + "#")
		} else {
			fmt.Println(strings.Repeat("#", in2))
		}
	}
}
