package main

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}

	L1 := len(a)
	L2 := len(b)
	if L1 < L2 {
		return L2
	}
	return L1
}
