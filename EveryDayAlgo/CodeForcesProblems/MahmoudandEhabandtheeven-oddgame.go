package main

import (
	"bufio"
	"fmt"
	"os"
)

func MahmoudandEhabandtheevenoddgame() {
	in := bufio.NewReader(os.Stdin)
	var number int
	fmt.Fscan(in, &number)
	if number%2 == 0 {
		fmt.Print("Mahmoud")
	} else {
		fmt.Print("Ehab")
	}
}
