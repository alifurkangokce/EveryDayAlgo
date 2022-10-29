package main

import "fmt"

func Hulk() {
	var inp int
	var result string
	fmt.Scan(&inp)
	for i := 1; i <= inp; i++ {
		if i == inp && (i%2 == 0) {
			result += "I love it "
		} else if i == inp && (i%2 != 0) {
			result += "I hate it "
		} else if i%2 == 0 {
			result += "I love that "
		} else {
			result += "I hate that "
		}
	}
	fmt.Print(result)
}
