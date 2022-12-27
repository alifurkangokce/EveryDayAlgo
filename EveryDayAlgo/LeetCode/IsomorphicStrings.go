package main

func isIsomorphic(s string, t string) bool {

	map1 := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if _, ok := map1[s[i]]; ok {
			if map1[s[i]] != t[i] {
				return false
			}
		} else {
			if checkValue(map1, t[i]) {
				return false
			}
			map1[s[i]] = t[i]
		}
	}
	return true
}

func checkValue(m map[byte]byte, val byte) bool {
	for _, x := range m {
		if x == val {
			return true
		}
	}
	return false
}
