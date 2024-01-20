package main

func restoreString(s string, indices []int) string {
	var res string
	arr := make([]byte, len(s))
	for k, v := range indices {
		arr[v] = s[k]
	}
	for _, v := range arr {
		res += string(v)
	}
	return res
}
