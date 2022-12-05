package main

import (
	"fmt"
	"sort"
)

func TheNewYearMeetingFriends() {
	var i1, i2, i3 int
	fmt.Scan(&i1, &i2, &i3)
	x := []int{i1, i2, i3}
	sort.Ints(x)
	fmt.Print((x[0]-x[1])*-1 + x[2] - x[1])
}
