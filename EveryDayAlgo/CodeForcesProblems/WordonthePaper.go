package main

import (
	"bufio"
	"fmt"
	"os"
)

func WordonthePaper() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var dot string
		var res string
		for i := 0; i < 8; i++ {
			fmt.Fscan(in, &dot)
			for _, v := range dot {
				if v != '.' {
					res += string(v)
				}
			}
		}
		fmt.Println(res)
	}
}
