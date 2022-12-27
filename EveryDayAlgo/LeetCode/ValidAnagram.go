package main

func isAnagram(s string, t string) bool {
	// if len(s) != len(t) {
	// 	return false
	// }
	// x := make(map[byte]int)
	// for i := 0; i <= len(s)-1; i++ {
	// 	x[s[i]]++
	// }
	// for i := 0; i <= len(t)-1; i++ {
	// 	x[t[i]]--
	// }
	// for a := 0; a <= len(s)-1; a++ {
	// 	if x[s[a]] > 0 {
	// 		return false
	// 	}
	// }
	// return true

	if len(s) != len(t) {
		return false
	}
	var m [26]int
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
		m[t[i]-'a']--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
