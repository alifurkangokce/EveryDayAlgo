package main

func interpret(command string) string {
	var res string
	for i := 0; i <= len(command)-1; i++ {
		if command[i] == '(' {
			if command[i+1] == ')' {
				res += "o"
				i += 1
			} else {
				res += "al"
				i += 3
			}
		} else {
			res += "G"
		}
	}
	return res
}
