package main

func toLowerCase(str string) string {
	s := make([]rune, len(str))
	i := 0
	for _, r := range str {
		if r >= 'A' && r <= 'Z' {
			s[i] = r + rune(32)
		} else {
			s[i] = r
		}

		i++
	}
	return string(s)
}
