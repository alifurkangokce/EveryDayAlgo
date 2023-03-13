package main

import (
	"bufio"
	"fmt"
	"os"
)

func FairPlayoff() {
	in := bufio.NewReader(os.Stdin)
	var t, n1, n2, n3, n4 int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n1, &n2, &n3, &n4)
		if ((n1 > n3 && n1 > n4) && (n2 > n3 && n2 > n4)) || ((n3 > n1 && n3 > n2) && (n4 > n1 && n4 > n2)) {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
