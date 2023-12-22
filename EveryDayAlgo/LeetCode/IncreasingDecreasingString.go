package main

import "strings"

type KeyValue struct {
	Key   int
	Value rune
}

func sortString(s string) string {

	maps := make(map[rune]int)
	for _, str := range s {
		maps[str]++
	}
	count := len(s)
	var sb strings.Builder
	for count != 0 {
		for i := 'a'; i <= 'z'; i++ {
			if maps[i] != 0 {
				maps[i]--
				count--
				sb.WriteRune(i)
			}
		}
		for i := 'z'; i >= 'a'; i-- {
			if maps[i] != 0 {
				maps[i]--
				count--
				sb.WriteRune(i)
			}
		}

	}

	return sb.String()
}
