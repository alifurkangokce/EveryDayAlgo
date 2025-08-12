package main

import (
	"bufio"
	"fmt"
	"os"
)

func XeniaAndRingroad() {
	in := bufio.NewReader(os.Stdin)

	var n int64
	var m int64
	if _, err := fmt.Fscan(in, &n, &m); err != nil {
		return
	}

	current := int64(1)
	var total int64 = 0

	for i := int64(0); i < m; i++ {
		var a int64
		fmt.Fscan(in, &a)

		if a >= current {
			total += a - current
		} else {
			// wrap around
			total += n - current + a
		}
		current = a
	}

	fmt.Println(total)
}
