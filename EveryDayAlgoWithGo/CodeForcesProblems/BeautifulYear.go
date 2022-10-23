package main

import "fmt"

func BeautifulYear() {
	var y int
	fmt.Scan(&y)

	for y++; len((map[int]bool{
		y % 10:       true,
		y / 10 % 10:  true,
		y / 100 % 10: true,
		y / 1000:     true})) != 4; y++ {
	}
	fmt.Println(y)
}
