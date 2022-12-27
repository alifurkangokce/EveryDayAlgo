package main

func countBits(num int) []int {
	res := make([]int, num+1)
	if num <= 0 {
		return res
	}

	for i := 0; (i << 1) <= num; i++ {
		if (i << 1) <= num {
			res[i<<1] = res[i]
		}

		if (i<<1)+1 <= num {
			res[(i<<1)+1] = res[i] + 1
		}
	}

	return res
}
