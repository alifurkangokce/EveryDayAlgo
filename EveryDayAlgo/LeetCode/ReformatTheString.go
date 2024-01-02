package main

import "sort"

func reformat(s string) string {
	var res string
	byte1 := []byte(s)

	sort.Slice(byte1, func(i, j int) bool {
		return byte1[i] < byte1[j]
	})
	i := 0
	for k, v := range byte1 {
		if v >= 'a' {
			i = k
			break
		}
	}
	if len(byte1[:i]) != len(byte1[i:]) && len(byte1[:i])+1 != len(byte1[i:]) && len(byte1[:i]) != len(byte1[i:])+1 {
		return ""
	}
	if len(byte1[:i]) == len(byte1[i:]) {
		j := 0
		for i := 0; i <= len(byte1)-1; i += 2 {
			res += string(byte1[j]) + string(byte1[len(byte1)-1-j])
			j++
		}
	} else if len(byte1[:i]) > len(byte1[i:]) {
		res += string(byte1[0])
		j := 0
		for i := 0; i < len(byte1)-1; i += 2 {
			res += string(byte1[len(byte1)-1-j]) + string(byte1[j+1])
			j++
		}
	} else {
		res += string(byte1[len(byte1)-1])
		j := 0
		for i := 0; i < len(byte1)-1; i += 2 {
			res += string(byte1[j]) + string(byte1[len(byte1)-2-j])
			j++
		}
	}

	return res
}
