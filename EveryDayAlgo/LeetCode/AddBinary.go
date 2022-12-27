package main

import "strconv"

func addBinary(a string, b string) string {
	var result string
	var carry int
	var i, j int
	for i < len(a) || j < len(b) {
		var sum int
		if i < len(a) {
			sum += int(a[len(a)-1-i] - '0')
		}
		if j < len(b) {
			sum += int(b[len(b)-1-j] - '0')
		}
		sum += carry
		carry = sum / 2
		result = strconv.Itoa(sum%2) + result
		i++
		j++
	}
	if carry > 0 {
		result = "1" + result
	}
	return result

}
