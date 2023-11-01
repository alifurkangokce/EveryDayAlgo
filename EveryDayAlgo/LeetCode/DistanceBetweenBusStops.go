package main

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if len(distance) == 1 || start == destination {
		return 0
	}
	if start > destination {
		start, destination = destination, start
	}
	t := sum(distance)
	a := sum(distance[start:destination])
	return min(a, t-a)
}

func sum(t []int) int {
	a := 0
	for _, v := range t {
		a += v
	}
	return a
}
