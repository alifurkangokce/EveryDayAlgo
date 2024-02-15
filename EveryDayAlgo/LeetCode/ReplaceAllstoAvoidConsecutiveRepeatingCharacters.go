package main

func modifyString(s string) string {
	if len(s) == 1 && s[0] == '?' {
		return "a"
	} else if len(s) == 1 {
		return s
	}
	bytes := []byte(s)
	for i := 0; i <= len(bytes)-1; i++ {
		if bytes[i] == '?' {
			str := byte('a')
			if i != 0 && i != len(bytes)-1 {
				for bytes[i+1] == str || bytes[i-1] == str {
					str++
				}
				bytes[i] = str
				i--
			} else {
				if i == 0 {
					for bytes[i+1] == str {
						str++
					}
				} else {
					for bytes[i-1] == str {
						str++
					}
				}

				bytes[i] = str
			}
		}
	}
	return string(bytes)
}
