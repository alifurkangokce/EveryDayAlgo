package main

import (
	"bufio"
	"fmt"
	"os"
)

func DoNotBeDistracted() {

	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		mm := make(map[byte]int)
		var n int
		var ll string
		yn := true
		fmt.Fscan(in, &n)
		fmt.Fscan(in, &ll)
		mm[ll[0]]++
		for i := 1; i < n; i++ {
			_, ok := mm[ll[i]]
			if ok {
				if ll[i-1] != ll[i] {
					fmt.Println("NO")
					yn = false
					break
				}
			} else {
				mm[ll[i]]++
			}

		}
		if yn {
			fmt.Println("YES")
		}

	}
}
