package main

import (
	"bufio"
	"fmt"
	"os"
)

func MediumNumber() {
	in := bufio.NewReader(os.Stdin)
	var t, n1, n2, n3 int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n1, &n2, &n3)
		if (n1 > n2 && n1 < n3) || (n1 < n2 && n1 > n3) {
			fmt.Println(n1)
		} else if (n2 > n1 && n2 < n3) || (n2 < n1 && n2 > n3) {
			fmt.Println(n2)
		} else if (n3 > n1 && n3 < n2) || (n3 < n1 && n3 > n2) {
			fmt.Println(n3)
		}
	}
}
