package main

func decrypt(code []int, k int) []int {
	ll := len(code)
	res := make([]int, ll)
	if k == 0 {
		return res
	} else if k > 0 {
		for i := 0; i < ll; i++ {
			sum := 0
			for j := 1; j <= k; j++ {
				sum += code[(i+j)%ll]
			}
			res[i] = sum
		}
	} else {
		for i := 0; i < ll; i++ {
			sum := 0
			for j := 1; j <= -k; j++ {
				sum += code[(i-j+ll)%ll]
			}
			res[i] = sum
		}
	}

	return res
}
