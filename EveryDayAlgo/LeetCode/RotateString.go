package main

import "strings"

func rotateString(s string, goal string) bool {
	// for i := 0; i < len(s)-1; i++ {
	// 	if s == goal {
	// 		return true
	// 	}
	// 	s = string(s[len(s)-1]) + string(s[0:len(s)-1])
	// }
	// return false

	return len(s) == len(goal) && strings.Contains(s+s, goal)
}
