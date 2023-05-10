package main

func backspaceCompare(s string, t string) bool {
	for i := 0; i <= len(s)-1; i++ {
		if s[i] == '#' {
			if i >= 2 {
				s = string(s[:i-1]) + s[i+1:]
				i = -1
			} else {
				s = s[i+1:]
				i = -1
			}
		}
	}
	for i := 0; i <= len(t)-1; i++ {
		if t[i] == '#' {
			if i >= 2 {
				t = string(t[:i-1]) + t[i+1:]
				i = -1
			} else {
				t = t[i+1:]
				i = -1
			}
		}
	}
	return s == t
}
