package main

import (
	"fmt"
)

func YoungPhysicist() {
	var xi, yi, zi int
	var x, y, z int
	var n int
	fmt.Scan(&n)
	for ; n > 0; n-- {
		fmt.Scan(&xi, &yi, &zi)
		x, y, z = x+xi, y+yi, z+zi
	}
	if x == 0 && y == 0 && z == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
