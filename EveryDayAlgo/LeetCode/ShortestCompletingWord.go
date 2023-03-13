package main

import (
	"math"
	"strings"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	licensePlate = strings.ToLower(licensePlate)
	mm := make(map[rune]int)
	res := ""
	ll := math.MaxInt
	for _, v := range licensePlate {
		if v >= 'a' && v <= 'z' {
			mm[v]++
		}
	}
	for _, v := range words {
		v = strings.ToLower(v)
		mmw := make(map[rune]int)
		active := true
		for _, w := range v {
			if w >= 'a' && w <= 'z' {
				mmw[w]++
			}
		}
		for k, v := range mm {
			if mmw[k] < v {
				active = false
				break
			}
		}
		if active {
			if len(v) < ll {
				res = v
				ll = len(v)
			}
		}
	}
	return res
}
