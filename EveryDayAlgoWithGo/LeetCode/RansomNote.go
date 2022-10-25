package main

func canConstruct(ransomNote string, magazine string) bool {
	chars := make([]int, 26)
	for _, c := range magazine {
		i := int(c - 'a')
		chars[i]++
	}

	for _, c := range ransomNote {
		i := int(c - 'a')
		chars[i]--
		if chars[i] < 0 {
			return false
		}
	}

	return true
}
