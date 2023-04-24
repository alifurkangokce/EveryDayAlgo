package main

import (
	"bufio"
	"fmt"
	"os"
)

func FafaAndHisCompany() {
	in := bufio.NewReader(os.Stdin)
	var inp, possibleCounter int
	fmt.Fscan(in, &inp)
	if inp <= 3 {
		fmt.Print(1)
		return
	} else {
		for i := 2; i <= inp/2; i++ {
			if inp%i == 0 {
				possibleCounter++
			}
		}
	}
	fmt.Print(possibleCounter + 1)
}
