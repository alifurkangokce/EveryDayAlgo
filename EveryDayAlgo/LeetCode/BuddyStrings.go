package main

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	diff := [2]int{}
	count := 0
	// store duplicate character
	dup := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			count++
			if count <= 2 {
				diff[count-1] = i
			} else {
				return false
			}
		}
		if _, ok := dup[s[i]]; ok {
			dup[s[i]]++
		} else {
			dup[s[i]] = 1
		}
	}
	// A equals to B
	if count == 0 {
		// check whether there has a duplicate character
		for _, v := range dup {
			if v > 1 {
				return true
			}
		}
	} else if count == 2 {
		if s[diff[0]] == goal[diff[1]] && s[diff[1]] == goal[diff[0]] {
			return true
		}
	}
	return false
}
