package main

import (
	"bufio"
	"fmt"
	"os"
)

func CardsForFriends() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)
	var w, h, n int
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &w, &h, &n)
		cnt := 1
		if n == 1 {
			fmt.Println("YES")
			continue
		}
		for w%2 == 0 {
			cnt *= 2
			w = w / 2
		}
		for h%2 == 0 {
			cnt *= 2
			h = h / 2
		}
		if cnt >= n {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
