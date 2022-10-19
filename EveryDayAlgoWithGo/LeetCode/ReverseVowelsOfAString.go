package main

func reverseVowels(s string) string {
	str := []byte(s)
	cnt := 1
	for i := 0; i <= len(s)-1; i++ {
		if isVovel(str[i]) {
			for j := len(s) - cnt; j > i; j-- {
				cnt++
				if isVovel(str[j]) {
					x := str[i]
					str[i] = str[j]
					str[j] = x
					break
				}
			}
		}
	}
	return string(str)
}

func isVovel(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' || ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U'
}
