package main

func maxLengthBetweenEqualCharacters(s string) int {

	var counter int = 1
	var check bool
	mm := make(map[rune]int)
	for _, v := range s {
		mm[v]++
	}
	for i := 0; i < len(s)-1; i++ {
		if mm[rune(s[i])] > 1 {
			for j := len(s) - 1; j >= i+1; j-- {
				if s[j] == s[i] {
					check = true
					if j-i > counter {
						counter = j - i
					}
					break
				}
			}
		}
	}
	if check && counter > 1 {
		return counter - 1
	} else if check {
		return 0
	}
	return -1

}
