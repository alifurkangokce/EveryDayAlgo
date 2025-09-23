package main

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	// var res string
	// var i int = 0
	// for i < len(word1) && i < len(word2) {
	// 	res += string(word1[i]) + string(word2[i])
	// 	i++
	// }
	// res += word1[i:] + word2[i:]
	// return res

	var b strings.Builder
	i := 0
	for i < len(word1) && i < len(word2) {
		b.WriteByte(word1[i])
		b.WriteByte(word2[i])
		i++
	}
	if i < len(word1) {
		b.WriteString(word1[i:])
	}
	if i < len(word2) {
		b.WriteString(word2[i:])
	}
	return b.String()
}
