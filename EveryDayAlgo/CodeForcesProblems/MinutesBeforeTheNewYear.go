package main

import (
	"bufio"
	"fmt"
	"os"
)

func MinutesBeforeTheNewYear() {
	in := bufio.NewReader(os.Stdin)
	var t, h, m int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &h, &m)
		fmt.Println((23-h)*60 + (60 - m))
	}
}
