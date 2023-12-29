package main

import (
	"bufio"
	"fmt"
	"os"
)

func TenWordsofWisdom() {
	in := bufio.NewReader(os.Stdin)
	var t, t2 int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &t2)
		var n1, n2, q, res int
		for i := 0; i < t2; i++ {
			fmt.Fscan(in, &n1, &n2)
			if n1 <= 10 && n2 > q {
				res = i + 1
				q = n2
			}
		}
		fmt.Println(res)
	}
}
