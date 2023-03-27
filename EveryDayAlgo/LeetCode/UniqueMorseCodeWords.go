package main

func uniqueMorseRepresentations(words []string) int {
	var ss []string = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	mm := make(map[string]string)
	for i := 0; i < 26; i++ {
		mm[string('a'+i)] = ss[i]
	}
	res := make(map[string]int)
	for _, v := range words {
		var str string
		for _, vv := range v {
			str += mm[string(vv)]
		}
		res[str]++
	}
	return len(res)
}
