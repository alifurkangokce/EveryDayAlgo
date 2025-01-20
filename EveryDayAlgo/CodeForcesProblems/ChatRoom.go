package main

import (
	"bufio"
	"fmt"
	"os"
)

func ChatRoom() {
	var inp, hello string = "", "hello"
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &inp)
	var cnt int
	for _, v := range inp {
		if hello[cnt] == byte(v) {
			cnt++
		}
		if cnt == 5 {
			fmt.Print("YES")
			return
		}
	}
	fmt.Print("NO")
}
