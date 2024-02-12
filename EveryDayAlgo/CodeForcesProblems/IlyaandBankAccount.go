package main

import (
	"fmt"
)

func IlyaandBankAccount() {
	var n int
	fmt.Scan(&n)
	if n >= 0 {
		fmt.Println(n)
	} else {
		k := n / 10
		j := n/100*10 + n%10
		if k > j {
			fmt.Println(k)
		} else {
			fmt.Println(j)
		}
	}
}
