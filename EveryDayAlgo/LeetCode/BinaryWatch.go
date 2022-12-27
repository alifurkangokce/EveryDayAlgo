package main

import "fmt"

func readBinaryWatch(num int) []string {

	res := []string{}

	for h := 0; h < 12; h++ {
		for m := 0; m < 60; m++ {
			if helper1(h)+helper1(m) == num {
				s := fmt.Sprintf("%d:%02d", h, m)
				res = append(res, s)
			}
		}
	}

	return res

}

func helper1(n int) int {
	res := 0
	for n > 0 {
		res += n % 2
		n /= 2
	}
	return res
}
