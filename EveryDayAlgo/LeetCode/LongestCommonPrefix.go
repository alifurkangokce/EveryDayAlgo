package main

import "strings"

func longestCommonPrefix(strs []string) string {
	result := strings.Builder{}
	if len(strs) == 0 {
		return result.String()
	}
	for i := 0; i < len(strs[0]); i++ {
		firstWord := strs[0][i]
		for _, s := range strs {
			if i >= len(s) || s[i] != firstWord {
				return result.String()
			}
		}

		result.WriteByte(firstWord)
	}

	return result.String()

}
