package main

func isOrderedPair(w1, w2 string, order map[byte]int) bool {
	for i := 0; i < len(w1) && i < len(w2); i++ {
		if order[w1[i]] != order[w2[i]] {
			return order[w1[i]] < order[w2[i]]
		}
	}
	return len(w1) <= len(w2)
}
func isAlienSorted(words []string, order string) bool {
	o := map[byte]int{}
	for i := 0; i < len(order); i++ {
		o[order[i]] = i
	}

	for i := 1; i < len(words); i++ {
		if !isOrderedPair(words[i-1], words[i], o) {
			return false
		}
	}
	return true
}
