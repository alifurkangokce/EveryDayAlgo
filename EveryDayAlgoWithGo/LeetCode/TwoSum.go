package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, number := range nums {
		if j, ok := m[target-number]; ok {
			return []int{j, index}
		}
		m[number] = index
	}
	return nil
}
func main() {
	fmt.Print(twoSum([]int{3, 2, 4}, 6))
}
