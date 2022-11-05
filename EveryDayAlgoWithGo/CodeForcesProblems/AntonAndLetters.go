package main

import (
	"bufio"
	"fmt"
	"os"
)

func AntonAndLetters() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')

	m := map[int32]bool{}

	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			m[c] = true
		}
	}

	fmt.Print(len(m))

}
