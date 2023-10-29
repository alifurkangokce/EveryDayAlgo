package main

func countCharacters(words []string, chars string) int {
	mm := make(map[rune]int)
	var res int
	for _, v := range chars {
		mm[v]++
	}
	for _, v := range words {
		if countCharHelper(mm, v) {
			res += len(v)
		}
	}
	return res
}
func countCharHelper(mm map[rune]int, v string) bool {
	ms := make(map[rune]int)
	for _, s := range v {
		ms[s]++
	}
	for k, v := range ms {
		if mm[k] < v {
			return false
		}
	}
	return true
}
