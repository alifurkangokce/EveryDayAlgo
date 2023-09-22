package main

import (
	"bufio"
	"fmt"
	"os"
)

func AtillasFavouriteProblem() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var n, max int
		var res rune
		fmt.Fscan(in, &n)
		var str string
		fmt.Fscan(in, &str)
		for i := 0; i <= n-1; i++ {
			res = 'a' - rune(str[i])
			if int(res)*-1 > max {
				max = int(res) * -1
			}
		}
		fmt.Println(max + 1)
	}
}
