package main

import "sort"

func mostVisited(n int, rounds []int) []int {
	ans := []int{}
	if len(rounds) == 0 {
		return ans
	}

	s := rounds[0]
	if rounds[len(rounds)-1] >= s {
		for i := s; i <= rounds[len(rounds)-1]; i++ {
			ans = append(ans, i)
		}
		return ans
	}
	for i := s; i < n+1; i++ {
		ans = append(ans, i)
	}
	for i := 1; i <= rounds[len(rounds)-1]; i++ {
		ans = append(ans, i)
	}
	sort.Ints(ans)
	return ans
}
