package main

import "strings"

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i := 0; i < len(words); i++ {
		words[i] = reverse(words[i])
	}
	return strings.Join(words, " ")
}

func reverse(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < int(n/2); i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}
