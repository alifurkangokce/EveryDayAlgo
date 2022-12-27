package main

import "strings"

func findWords(words []string) []string {
	k1S := "qwertyuiop"
	k2S := "asdfghjkl"
	k3S := "zxcvbnm"
	var result []string
	for _, v1 := range words {
		v := strings.ToLower(v1)
		ch1, ch2, ch3 := true, true, true
		if strings.Contains(k1S, string(v[0])) {
			for i := 1; i <= len(v)-1; i++ {
				if !strings.Contains(k1S, string(v[i])) {
					ch1 = false
					break
				}
			}
		} else if strings.Contains(k2S, string(v[0])) {
			for i := 1; i <= len(v)-1; i++ {
				if !strings.Contains(k2S, string(v[i])) {
					ch2 = false
					break
				}
			}

		} else if strings.Contains(k3S, string(v[0])) {
			for i := 1; i <= len(v)-1; i++ {
				if !strings.Contains(k3S, string(v[i])) {
					ch3 = false
					break
				}
			}
		}
		if ch1 && ch2 && ch3 {
			result = append(result, v1)
		}
	}
	return result
}
