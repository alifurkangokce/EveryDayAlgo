package main

import (
	"bytes"
	"strings"
)

func licenseKeyFormatting(s string, k int) string {
	s = strings.ReplaceAll(s, "-", "")
	data := make([][]byte, (len(s)-1)/k+1)
	sByte := []byte(s)
	i := len(s)
	for j := len(data) - 1; i > k; i -= k {
		data[j] = bytes.ToUpper(sByte[i-k : i])
		j--
	}
	data[0] = bytes.ToUpper(sByte[:i])
	return string(bytes.Join(data, []byte("-")))
}
