package main

import (
	"bufio"
	"fmt"
	"os"
)

func Candies() {
	// Fast input/output
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// Read the number of test cases
	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(reader, &n)

		// Iterate over potential values of k
		for k := 2; ; k++ {
			// Calculate 2^k - 1
			sum := (1 << k) - 1

			// Check if n is divisible by sum
			if n%sum == 0 {
				// If divisible, calculate x
				x := n / sum
				fmt.Fprintln(writer, x)
				break
			}
		}
	}
}
