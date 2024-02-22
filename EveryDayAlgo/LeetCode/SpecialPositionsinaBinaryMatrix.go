package main

func numSpecial(M [][]int) int {
	nx, ny := len(M[0]), len(M)
	data_x := make([]int, nx)
	data_y := make([]int, ny)

	for y := 0; y < ny; y++ {
		for x := 0; x < nx; x++ {
			if M[y][x] == 1 {
				data_x[x]++
				data_y[y]++
			}
		}
	}

	res := 0
	for y := 0; y < ny; y++ {
		for x := 0; x < nx; x++ {
			if M[y][x] == 1 && data_x[x] == 1 && data_y[y] == 1 {
				res++
			}
		}
	}

	return res
}
