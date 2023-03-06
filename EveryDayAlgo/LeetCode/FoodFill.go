package main

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}

	floodFillHelper(image, sr, sc, image[sr][sc], newColor)
	return image
}

var directionsHelper = [][]int{
	[]int{0, 1},
	[]int{0, -1},
	[]int{1, 0},
	[]int{-1, 0},
}

func floodFillHelper(image [][]int, i, j int, oldColor, newColor int) {
	if i < 0 || i >= len(image) || j < 0 || j >= len(image[0]) || image[i][j] != oldColor {
		return
	}

	image[i][j] = newColor
	for _, direction := range directionsHelper {
		ii, jj := i+direction[0], j+direction[1]
		floodFillHelper(image, ii, jj, oldColor, newColor)
	}
}
