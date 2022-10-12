package main

import (
	"fmt"
	"sort"
	"strings"
)

func HelpFulMath() {
	var i string
	fmt.Scan(&i)
	sprint := strings.Split(i, "+")
	sort.Strings(sprint)
	var res string
	for i := 0; i <= len(sprint)-1; i++ {
		res += sprint[i] + "+"
	}
	fmt.Print(strings.TrimRight(res, "+"))
}
