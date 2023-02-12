package main

import (
	"bufio"
	"fmt"
	"os"
)

func FairDivision() {
	in := bufio.NewReader(os.Stdin)
	var t, t2, n int
	mm := make(map[int]int)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &t2)
		for j := 0; j < t2; j++ {
			fmt.Fscan(in, &n)
			mm[n]++
		}
		if mm[1] == 0 && mm[2]%2 == 0 {
			fmt.Println("YES")
		} else if mm[1]%2 == 0 && mm[1] != 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		mm[1] = 0
		mm[2] = 0
	}

}
