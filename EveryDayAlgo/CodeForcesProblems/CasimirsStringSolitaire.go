package main

import (
	"bufio"
	"fmt"
	"os"
)

func CasimirsStringSolitaire() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)
	var s string
	for i := 0; i < t; i++ {
		mm := make(map[byte]int)
		fmt.Fscan(in, &s)
		for i := 0; i <= len(s)-1; i++ {
			mm[s[i]]++
		}
		if mm['A']+mm['C'] == mm['B'] {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
