package main

func numberOfLines(widths []int, s string) []int {
	mm := make(map[string]int)
	lineCount := 1
	sT := 0
	for i := 0; i < 26; i++ {
		mm[string('a'+i)] = widths[i]
	}
	for i := 0; i <= len(s)-1; i++ {
		sT += mm[string(s[i])]
		if sT > 100 {
			sT = 0
			lineCount++
			sT += mm[string(s[i])]
		}
	}
	return []int{lineCount, sT}
}
