package main

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var str1, str2 string
	for _, v := range word1 {
		str1 += v
	}
	for _, v := range word2 {
		str2 += v
	}
	return str1 == str2
}
