package main

import "fmt"

func QueueAtTheSchool() {
	var inp, sec int
	var chld string
	fmt.Scan(&inp, &sec, &chld)
	str := []byte(chld)
	for j := 0; j < sec; j++ {
		for i := 0; i < inp-1; i++ {
			if str[i] == 'B' && str[i+1] == 'G' {
				x := str[i]
				str[i] = str[i+1]
				str[i+1] = x
				i++
			}
		}
	}
	fmt.Print(string(str))
}
