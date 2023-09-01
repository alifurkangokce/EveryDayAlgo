package main

import (
	"fmt"
)

func Colourblindness() {
	var t int
	fmt.Scan(&t)
	for j := 0; j < t; j++ {
		var n int
		fmt.Scan(&n)
		var s1, s2 string
		fmt.Scan(&s1)
		fmt.Scan(&s2)
		flag := 1
		for i := 0; i < n; i++ {
			if s1[i] != s2[i] {
				if (s1[i] == 'B' && s2[i] == 'G') || (s1[i] == 'G' && s2[i] == 'B') {
					continue
				} else {
					flag = 0
					break
				}
			}
		}
		if flag == 1 {
			fmt.Printf("YES\n")
		} else {
			fmt.Printf("NO\n")
		}
	}
}
