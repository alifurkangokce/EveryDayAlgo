package main

import "strconv"

func thousandSeparator(n int) string {
	strInt := strconv.Itoa(n)
	ll := len(strInt)
	cnt := ll / 3
	mod := ll % 3
	var res string
	if cnt > 0 && ll > 3 {
		if mod > 0 {
			res += strInt[0:mod] + "."
		}
		for i := 0; i < cnt-1; i++ {
			res += strInt[mod:mod+3] + "."
			mod += 3
		}
		res += strInt[mod : mod+3]
		return res
	} else {
		return strInt
	}

}
