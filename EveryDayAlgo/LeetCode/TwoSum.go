package main

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, number := range nums {
		if j, ok := m[target-number]; ok {
			return []int{j, index}
		}
		m[number] = index
	}
	return nil
}
