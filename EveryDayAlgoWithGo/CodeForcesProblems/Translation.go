package main

import "fmt"

func Translation() {
	var str1, str2 string
	fmt.Scan(&str1, &str2)
	ln := len(str1)
	for i := 0; i <= ln-1; i++ {
		if str1[i] != str2[ln-1-i] {
			fmt.Print("NO")
			return
		}
	}
	fmt.Print("YES")
}
