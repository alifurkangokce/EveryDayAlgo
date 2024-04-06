package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StringTask() {
	in := bufio.NewReader(os.Stdin)
	var inp string
	fmt.Fscan(in, &inp)
	var vovels string = "aeiouy"
	inp = strings.ToLower(inp)
	var res string
	for i := 0; i <= len(inp)-1; i++ {
		if !strings.Contains(vovels, string(inp[i])) {
			res += "." + string(inp[i])
		}
	}
	fmt.Println(res)
}
