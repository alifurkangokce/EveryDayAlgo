package main

import (
	"fmt"
	"sort"
)

func Expression() {
	var a, b, c int
	fmt.Scanf("%d\n%d\n%d\n", &a, &b, &c)
	ans := []int{(a + b) * c, a * (b + c), a + b + c, a * b * c, a + b*c, a*b + c}
	sort.Ints(ans)
	fmt.Print(ans[len(ans)-1])
}
