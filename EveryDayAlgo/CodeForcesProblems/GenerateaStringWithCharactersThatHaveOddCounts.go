package main

import "strings"

func generateTheString(n int) string {
	return "b" + strings.Repeat(string('a'+n%2), n-1)
}
