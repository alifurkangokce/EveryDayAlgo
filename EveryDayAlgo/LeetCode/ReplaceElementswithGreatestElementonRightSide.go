package main

func replaceElements(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		arr[i] = getMaxElementsInArr(arr[i+1:])
	}
	arr[len(arr)-1] = -1
	return arr
}

func getMaxElementsInArr(arr []int) int {
	var max int = 0
	for _, v := range arr {
		if v >= max {
			max = v
		}
	}
	return max
}
