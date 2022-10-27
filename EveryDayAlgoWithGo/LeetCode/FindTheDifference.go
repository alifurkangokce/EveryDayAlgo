package main

func findTheDifference(s string, t string) byte {
	// x := make(map[byte]int)
	// y := make(map[byte]int)

	// for i := 0; i <= len(s)-1; i++ {
	// 	x[s[i]]++
	// }
	// for i := 0; i <= len(t)-1; i++ {
	// 	y[t[i]]++
	// }
	// for i := 0; i <= len(t)-1; i++ {
	// 	if x[t[i]] != y[t[i]] {
	// 		return t[i]
	// 	}
	// }
	// return '"'

	var diff1, diff2 int = int(t[0]), 0
	for i := 1; i < len(t); i++ {
		diff1 += int(t[i])
		diff2 += int(s[i-1])
	}
	return byte(diff1 - diff2)
}
