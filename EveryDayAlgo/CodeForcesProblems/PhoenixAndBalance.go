package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func PhoenixAndBalance() {
	in := bufio.NewReader(os.Stdin)
	var t, n int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var n1, n2 int
		fmt.Fscan(in, &n)
		for i := 1; i <= (n/2)-1; i++ {
			n1 += int(math.Pow(2, float64(i)))
		}
		n1 += int(math.Pow(2, float64(n)))
		for i := (n / 2); i < n; i++ {
			n2 += int(math.Pow(2, float64(i)))
		}
		fmt.Println(int32(math.Abs(float64(n1 - n2))))
	}
}
