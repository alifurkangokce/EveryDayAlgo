package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GennadyandaCardGame() {
	in := bufio.NewReader(os.Stdin)
	var s1, s2 string
	var check bool
	fmt.Fscan(in, &s1)
	for i := 0; i < 5; i++ {
		fmt.Fscan(in, &s2)
		if !check {
			if strings.Contains(s1, string(s2[0])) || strings.Contains(s1, string(s2[1])) {
				check = true
			}
		}
	}
	if check {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}

}
