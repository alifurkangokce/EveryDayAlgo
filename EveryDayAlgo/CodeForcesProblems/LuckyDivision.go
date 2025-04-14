package main

import (
	"bufio"
	"fmt"
	"os"
)

func LuckyDivision() {
	in := bufio.NewReader(os.Stdin)
	var s int
	fmt.Fscan(in, &s)
	if s == 4 || s == 47 || s == 7 || s == 774 || s == 477 || s%4 == 0 || s%7 == 0 || s%47 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
