package main

func isOneBitCharacter(bits []int) bool {
	p := 0
	for p < len(bits)-1 {
		if bits[p] == 1 {
			p += 2
		} else {
			p += 1
		}
	}
	return p == len(bits)-1
}
