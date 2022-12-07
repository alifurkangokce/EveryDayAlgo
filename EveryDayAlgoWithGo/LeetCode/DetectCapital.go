package main

func detectCapitalUse(word string) bool {
	if len(word) == 2 {
		return true
	}
	if word[0] >= 65 && word[0] <= 90 {
		if word[1] >= 65 && word[1] <= 90 {
			for i := 2; i <= len(word)-1; i++ {
				if word[i] < 65 || word[i] > 90 {
					return false
				}
			}
		} else {
			for i := 1; i <= len(word)-1; i++ {
				if word[i] < 'a' {
					return false
				}
			}
		}

	} else {
		for i := 1; i <= len(word)-1; i++ {
			if word[i] < 'a' {
				return false
			}
		}
	}
	return true
}
