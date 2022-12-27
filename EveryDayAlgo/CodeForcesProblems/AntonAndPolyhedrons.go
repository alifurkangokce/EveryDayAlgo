package main

import (
	"bufio"
	"fmt"
	"os"
)

func AntonAndPolyhedrons() {
	x := map[string]int{
		"Tetrahedron":  4,
		"Cube":         6,
		"Octahedron":   8,
		"Dodecahedron": 12,
		"Icosahedron":  20,
	}
	var res int
	var n int
	fmt.Scanf("%d \n", &n)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		scanner.Scan()
		s := scanner.Text()
		res += x[s]
	}
	fmt.Print(res)
}
