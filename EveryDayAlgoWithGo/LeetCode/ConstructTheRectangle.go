package main

func constructRectangle(area int) []int {
	l, w := area, 1
	for i := 1; i*i <= area; i++ {
		if area%i == 0 {
			l = i
			w = area / i
		}
	}
	if w > l {
		l, w = w, l
	}
	return []int{l, w}
}
