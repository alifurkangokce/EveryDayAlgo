package main

import "sort"

func lastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}

	for len(stones) != 1 || len(stones) == 0 {
		ls := len(stones)
		sort.Ints(stones)
		difstones := stones[ls-1] - stones[ls-2]
		if difstones != 0 {
			stones = stones[:ls-2]
			stones = append(stones, difstones)
		} else {
			stones = stones[:ls-2]
		}
		if len(stones) == 0 {
			return 0
		}

	}

	return stones[0]

}
