package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CapsLock() {

	var str string
	var check bool
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &str)
	check = false
	for i, v2 := range str {
		if v2 >= 97 && i != 0 {
			check = true
			break
		}
	}
	if check {
		fmt.Print(str)
	} else if str[0] < 97 {
		fmt.Print(strings.ToLower(str))
	} else {
		fmt.Print(strings.ToUpper(string(str[0])) + strings.ToLower(string(str[1:])))
	}

}
