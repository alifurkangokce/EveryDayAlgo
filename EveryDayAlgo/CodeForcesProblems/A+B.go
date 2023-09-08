package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AB() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)
	var s string
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &s)
		sSplit := strings.Split(s, "+")
		i1, _ := strconv.Atoi(sSplit[0])
		i2, _ := strconv.Atoi(sSplit[1])
		fmt.Println(i1 + i2)
	}
}
