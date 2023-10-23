package main

func duplicateZeros(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			x := arr[i+1:]
			y := arr[i:]
			copy(x, y)
			i++
		}
	}
}
