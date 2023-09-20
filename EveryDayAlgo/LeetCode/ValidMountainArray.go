package main

func validMountainArray(arr []int) bool {
	if len(arr) == 1 {
		return false
	}
	if arr[0] > arr[1] {
		return false
	}
	var increase bool = true
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			increase = false

		} else if arr[i] < arr[i+1] {
			if !increase {
				return false
			}
		} else {
			return false
		}
	}
	if increase {
		return false
	}
	return true
}
