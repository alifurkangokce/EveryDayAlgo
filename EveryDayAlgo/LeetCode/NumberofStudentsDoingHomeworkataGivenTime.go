package main

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	var res int
	for i := 0; i <= len(startTime)-1; i++ {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			res++
		}
	}
	return res
}
