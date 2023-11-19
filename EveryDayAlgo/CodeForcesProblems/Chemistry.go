package main

import "fmt"

func Chemistry() {
	var t int
	fmt.Scan(&t)
	for t > 0 {
		var n, k int

		fmt.Scan(&n, &k)
		var str string
		fmt.Scan(&str)
		byte1 := []byte(str)
		mp := make(map[byte]int)
		for _, b := range byte1 {
			mp[b]++
		}
		var ans int
		for _, m := range mp {
			if m%2 == 1 {
				ans++
			}
		}
		if ans <= k+1 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		t--
	}

}
