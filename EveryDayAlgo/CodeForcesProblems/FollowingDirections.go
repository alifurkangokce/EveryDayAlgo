package main

import (
	"bufio"
	"fmt"
	"os"
)

func FollowingDirection() {
	var t int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	var n int
	for i := 0; i < t; i++ {
		isYes := false
		fmt.Fscan(in, &n)
		var s string
		fmt.Fscan(in, &s)
		var x, y int
		for i := 0; i <= n-1; i++ {
			if s[i] == 'U' {
				x++
			} else if s[i] == 'D' {
				x--
			} else if s[i] == 'L' {
				y--
			} else {
				y++
			}
			if x == 1 && y == 1 {
				isYes = true
				break
			}

		}
		if isYes {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
