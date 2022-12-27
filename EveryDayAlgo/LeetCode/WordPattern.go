package main

import "strings"

func wordPattern(pattern string, s string) bool {
	maps := make(map[string]string)
	maps2 := make(map[string]string)
	splits := strings.Split(s, " ")
	if len(splits) != len(pattern) {
		return false
	}
	for i := 0; i <= len(pattern)-1; i++ {
		value, ok := maps[string(pattern[i])]
		value2, ok2 := maps2[string(splits[i])]
		if !ok && !ok2 {
			maps[string(pattern[i])] = string(splits[i])
			maps2[string(splits[i])] = string(pattern[i])
		} else {
			if value == splits[i] && value2 == string(pattern[i]) {
				continue
			} else {
				return false
			}
		}

	}
	return true
}
