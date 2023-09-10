package main

func isLetter(b byte) bool {
	return b >= 'A' && b <= 'Z' || b >= 'a' && b <= 'z'
}

func reverseOnlyLetters(S string) string {
	buf := []byte(S)
	l, h := 0, len(buf)-1

	for l < h {
		if !isLetter(buf[l]) { // skip non-letter on the left
			l += 1
		} else if !isLetter(buf[h]) { // skip non-letter on the right
			h -= 1
		} else { // swap letters and move pointers
			buf[l], buf[h] = buf[h], buf[l]
			l += 1
			h -= 1
		}
	}
	return string(buf)
}
