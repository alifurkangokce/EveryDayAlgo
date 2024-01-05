package main

func destCity(paths [][]string) string {
	mm := make(map[string]string)
	for _, v := range paths {
		mm[v[0]] = v[1]
	}
	for _, v := range mm {
		_, exist := mm[v]
		if !exist {
			return v
		}
	}
	return ""
}
