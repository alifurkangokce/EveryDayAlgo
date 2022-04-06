package main

// func lengthOfLastWord(s string) int {
// 	x := strings.Split(s, " ")
// 	leng := len(x) - 1
// 	for x[leng] == "" {
// 		leng--
// 	}
// 	return len(x[leng])
// }

func lengthOfLastWord(s string) int {
	count := 0
	for i := len(s) - 1; i > -1; i-- {
		if string(s[i]) != " " {
			count++
		} else if count > 0 {
			return count
		}
	}
	return count
}
