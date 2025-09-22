package main

import "unicode"

func longestNiceSubstring(s string) string {
	if len(s) < 2 {
		return ""
	}

	// karakterleri set içine koy
	charSet := make(map[rune]bool)
	for _, ch := range s {
		charSet[ch] = true
	}

	// her karakteri kontrol et
	for i, ch := range s {
		if !charSet[swapCase(ch)] {
			// bu karakter nice değil → böl
			left := longestNiceSubstring(s[:i])
			right := longestNiceSubstring(s[i+1:])
			if len(left) >= len(right) {
				return left
			}
			return right
		}
	}

	// bütün karakterler dengeli → nice string
	return s
}

// küçük ↔ büyük dönüşüm
func swapCase(ch rune) rune {
	if unicode.IsLower(ch) {
		return unicode.ToUpper(ch)
	}
	return unicode.ToLower(ch)
}
