package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Dubstep() {
	in := bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscan(in, &s)
	res := strings.Split(s, "WUB")
	for _, v := range res {
		if v != "" {
			fmt.Printf("%s ", v)
		}

	}
}
