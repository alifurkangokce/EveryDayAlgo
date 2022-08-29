package main

func titleToNumber(columnTitle string) int {
	rs := []rune(columnTitle)
	ret := 0
	for i := 0; i < len(rs); i++ {
		ret = ret*26 + (int(rs[i]-'A') + 1)
	}
	return ret
}
