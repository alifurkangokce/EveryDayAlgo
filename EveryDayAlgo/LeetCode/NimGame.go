package main

func canWinNim(n int) bool {
	if n > 3 {
		if n%4 == 0 {
			return false
		}
		return true
	}
	return true
}
