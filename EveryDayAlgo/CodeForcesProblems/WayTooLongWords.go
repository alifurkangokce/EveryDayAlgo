package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func wayTooLong() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)
	var s string
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &s)
		if len(s) <= 10 {
			fmt.Println(s)
		} else {
			res := string(s[0]) + strconv.Itoa(len(s)-2) + string(s[len(s)-1])
			fmt.Println(res)
		}
	}
}
