package main

func findPoisonedDuration(timeSeries []int, duration int) int {
	var result int
	for i := 0; i < len(timeSeries)-1; i++ {
		if timeSeries[i+1]-timeSeries[i] < duration {
			result += timeSeries[i+1] - timeSeries[i]
		} else {
			result += duration
		}
	}
	return result + duration
}
