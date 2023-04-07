package main

import "strings"

func toGoatLatin(sentence string) string {
	var output string
	var ma string = "maa"
	oS := strings.Split(sentence, " ")
	for _, v := range oS {
		if v[0] == 'a' || v[0] == 'e' || v[0] == 'i' || v[0] == 'o' || v[0] == 'u' || len(v) == 1 || v[0] == 'A' || v[0] == 'E' || v[0] == 'I' || v[0] == 'O' || v[0] == 'U' {
			output += v + ma + " "
		} else {
			output += string(v[1:]) + string(v[0]) + ma + " "
		}
		ma += "a"
	}
	return output[:len(output)-1]
}
