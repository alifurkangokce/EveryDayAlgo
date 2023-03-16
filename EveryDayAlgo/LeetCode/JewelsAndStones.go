package main

func numJewelsInStones(jewels string, stones string) int {
	mStones := make(map[rune]int)
	var res int
	for _, v := range stones {
		mStones[v]++
	}
	for _, v := range jewels {
		if mStones[v] > 0 {
			res += mStones[v]
		}
	}
	return res
}
