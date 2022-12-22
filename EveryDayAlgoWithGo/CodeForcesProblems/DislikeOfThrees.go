package main

import "fmt"

func DislikeOfThrees() {
	var (
		t int
		k int
	)
	fmt.Scanf("%d\n", &t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d\n", &k)
		var (
			t int
			n int
		)
		t = 1
		for j := 1; n != k; j++ {
			if j%3 == 0 || j%10 == 3 {
				j++
			}
			if j%3 == 0 || j%10 == 3 {
				j++
			}
			t = j
			n++
		}
		fmt.Println(t)
	}
}
