package main

import (
	"bufio"
	"fmt"
	"os"
)

func GrabtheCandies() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var n, number int
		var arr1, arr2 int
		fmt.Fscan(in, &n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &number)
			if number%2 == 0 {
				arr1 += number
			} else {
				arr2 += number
			}
		}
		if arr1 > arr2 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
