package main

import (
	"bufio"
	"fmt"
	"os"
)

func LoveStory() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var s string
		var codeforces = "codeforces"
		var res int
		fmt.Fscan(in, &s)
		for i := 0; i <= len(codeforces)-1; i++ {
			if s[i] != codeforces[i] {
				res++
			}
		}
		fmt.Println(res)
	}
}
