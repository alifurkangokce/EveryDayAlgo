package main

func judgeCircle(moves string) bool {
	mm := make(map[byte]int)
	for i := 0; i <= len(moves)-1; i++ {
		mm[moves[i]]++
	}
	return mm['U'] == mm['D'] && mm['L'] == mm['R']
}
