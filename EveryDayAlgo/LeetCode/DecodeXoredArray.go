package main

func decode(encoded []int, first int) []int {
	var res []int
	res = append(res, first)
	for i := 0; i <= len(encoded)-1; i++ {
		res = append(res, res[i]^encoded[i])
	}
	return res
}
