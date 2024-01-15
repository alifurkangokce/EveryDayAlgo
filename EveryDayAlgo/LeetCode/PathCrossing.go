package main

import "fmt"

func isPathCrossing(path string) bool {
	visitedPoints := map[string]bool{}
	currentX, currentY := 0, 0
	visitedPoints[fmt.Sprintf("%d,%d", currentX, currentY)] = true

	for _, direction := range path {
		switch direction {
		case 'N':
			currentY++
		case 'S':
			currentY--
		case 'E':
			currentX++
		case 'W':
			currentX--
		}

		point := fmt.Sprintf("%d,%d", currentX, currentY)
		if visitedPoints[point] {
			return true
		}

		visitedPoints[point] = true
	}

	return false
}
