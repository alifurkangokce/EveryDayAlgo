package main

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	splitted := strings.Split(sentence, " ")
	for k1, v := range splitted {
		var cnt int
		if len(v) >= len(searchWord) {
			for k, w := range searchWord {
				if w == rune(v[k]) {
					cnt++
				} else if len(v[k:]) < len(searchWord) {
					break
				}
			}
		}
		if cnt == len(searchWord) {
			return k1 + 1
		}
	}
	return -1
}
