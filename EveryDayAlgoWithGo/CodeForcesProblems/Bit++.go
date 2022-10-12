package main

import "fmt"

func BitPlus() {
	var input int
	var arrString [150]string
	var result int
	fmt.Scan(&input)
	for i := 0; i < input; i++ {
		var inpStr string
		fmt.Scan(&inpStr)
		arrString[i] = inpStr
	}
	for i := 0; i <= len(arrString)-1; i++ {
		if arrString[i] == "++X" || arrString[i] == "X++" {
			result++
		}
		if arrString[i] == "--X" || arrString[i] == "X--" {
			result--
		}
	}
	fmt.Print(result)
}
