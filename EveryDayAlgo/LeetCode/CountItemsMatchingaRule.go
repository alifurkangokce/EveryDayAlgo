package main

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	var res int
	var k int = 0
	if ruleKey == "type" {
		k = 0
	} else if ruleKey == "color" {
		k = 1
	} else {
		k = 2
	}
	for i := 0; i <= len(items)-1; i++ {
		if items[i][k] == ruleValue {
			res++
		}
	}
	return res
}
