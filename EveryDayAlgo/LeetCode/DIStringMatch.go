package main

func diStringMatch(s string) []int {
	var res []int
	var inc, dec int = 0, len(s)
	for i := 0; i <= len(s)-1; i++ {
		if s[i] == 'I' {
			res = append(res, inc)
			inc++
		} else {
			res = append(res, dec)
			dec--
		}
	}
	res = append(res, inc)
	return res
}
