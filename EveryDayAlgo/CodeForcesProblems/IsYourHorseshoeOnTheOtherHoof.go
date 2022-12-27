package main

import "fmt"

func IsYourHorsesShoe() {
	var i1, i2, i3, i4 int
	result := 3
	fmt.Scan(&i1, &i2, &i3, &i4)
	if i1 != i2 && i1 != i3 && i1 != i4 {
		result--
	}
	if i2 != i3 && i2 != i4 {
		result--
	}
	if i3 != i4 {
		result--
	}
	fmt.Print(result)
}
