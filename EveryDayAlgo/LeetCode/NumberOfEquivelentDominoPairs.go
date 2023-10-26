package main

func numEquivDominoPairs(dominoes [][]int) int {
	var cnt int
	for i := 0; i < len(dominoes)-1; i++ {
		for j := i + 1; j <= len(dominoes)-1; j++ {
			a := dominoes[i][0]
			b := dominoes[i][1]
			c := dominoes[j][0]
			d := dominoes[j][1]
			if (a == c && b == d) || (a == d && b == c) {
				cnt++
			}
		}

	}
	return cnt
}
