package main

import (
	"fmt"
	"strings"
)

func PetyaStrings() {
	var i1 string
	var i2 string
	fmt.Scan(&i1, &i2)
	i1 = strings.ToLower(i1)
	i2 = strings.ToLower(i2)
	if i1 == i2 {
		fmt.Print(0)
		return
	}
	for i := 0; i <= len(i1)-1; i++ {
		if i1[i] != i2[i] {
			if i1[i] > i2[i] {
				fmt.Print(1)
				return
			} else {
				fmt.Print(-1)
				return
			}
		}
	}
}
