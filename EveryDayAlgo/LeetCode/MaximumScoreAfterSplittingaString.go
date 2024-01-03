package main

func maxScore(s string) int {
	var max int
	for i := 1; i <= len(s)-1; i++ {
		z, o := maxScoreHepler(s[:i], s[i:])
		if z+o > max {
			max = z + o
		}
	}
	return max
}
func maxScoreHepler(z string, zz string) (int, int) {
	var zero, one int
	for _, v := range z {
		if v == '0' {
			zero++
		}
	}
	for _, v := range zz {
		if v == '1' {
			one++
		}
	}
	return zero, one
}
