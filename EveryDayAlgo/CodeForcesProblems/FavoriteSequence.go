package main

import (
	"bufio"
	"fmt"
	"os"
)

func FavoriteSequence() {
	in := bufio.NewReader(os.Stdin)
	var t, n int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var number int
		var arr []int
		fmt.Fscan(in, &n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &number)
			arr = append(arr, number)
		}
		for i := 0; i < (len(arr) / 2); i++ {

			fmt.Printf("%d ", arr[i])
			fmt.Printf("%d ", arr[len(arr)-1-i])
		}
		if len(arr)%2 != 0 {
			fmt.Printf("%d ", arr[len(arr)/2])

		}

	}
}
