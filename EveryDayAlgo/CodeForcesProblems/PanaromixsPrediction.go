package main

import "fmt"

func PanaromixsPrediction() {
	var t1, t2 int
	fmt.Scan(&t1, &t2)
	var check bool = false
	for i := t1 + 1; i <= t2; i++ {
		if PanaromixHelper(i) {
			check = true
			if i == t2 {
				fmt.Print("YES")
				break
			} else {
				fmt.Print("NO")
				break
			}
		}
	}
	if !check {
		fmt.Print("NO")
	}

}
func PanaromixHelper(t int) bool {
	for i := 2; i <= t/2; i++ {
		if t%i == 0 {
			return false
		}
	}
	return true
}
