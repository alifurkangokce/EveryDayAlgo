package main

func commonChars(words []string) []string {
	ff := make(map[rune]int)
	var res []string
	for _, v := range words[0] {
		ff[v]++
	}
	for _, v := range words[1:] {
		mm := make(map[rune]int)
		for _, k := range v {
			mm[k]++
		}
		for k, v := range ff {
			if mm[k] == 0 {
				delete(ff, k)
			} else {
				if mm[k] < v {
					ff[k] = mm[k]
				}
			}
		}
	}
	for k, v := range ff {
		for i := 0; i < v; i++ {
			res = append(res, string(k))
		}

	}
	return res
}
