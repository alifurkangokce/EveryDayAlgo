package main

func reorderSpaces(text string) string {
	var str []string
	var counter, spaces, modSpaces int
	var word, result string
	for _, v := range text {
		if v == ' ' {
			counter++
			if word != "" {
				str = append(str, word)
			}
			word = ""
		} else {
			word += string(v)
		}

	}
	if word != "" {
		str = append(str, word)
	}
	if len(str) == 1 {
		spaces = 0
		modSpaces = counter
	} else {
		spaces = counter / (len(str) - 1)
		modSpaces = counter % (len(str) - 1)
	}

	for i := 0; i < len(str)-1; i++ {
		result += str[i]
		for j := 0; j < spaces; j++ {
			result += " "
		}
	}
	result += str[len(str)-1]
	if modSpaces > 0 {

		for j := 0; j < modSpaces; j++ {
			result += " "
		}
	}
	return result

}
