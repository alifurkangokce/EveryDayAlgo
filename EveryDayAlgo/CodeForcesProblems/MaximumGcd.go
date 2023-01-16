package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func MaxiumumGcd() {
	in := bufio.NewReader(os.Stdin)
	var t, n int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		fmt.Println(math.Round(float64(n-1) / float64(2)))
	}
}
