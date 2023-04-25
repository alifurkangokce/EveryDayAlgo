package main

import (
	"fmt"
	"math"
)

func NightAtTheMuseum() {
	var s string
	fmt.Scanf("%s", &s)

	current := 97
	answer := 0

	for i := range s {
		difference := int(math.Abs(float64(current - int(s[i]))))
		if difference < 13 {
			answer += difference
		} else {
			answer += 26 - difference
		}
		current = int(s[i])
	}

	fmt.Print(answer)
}
