package main

func findDisappearedNumbers(nums []int) []int {
	freq := make([]int, len(nums)+1)

	for _, v := range nums {
		freq[v]++
	}

	var ans []int
	for i := 1; i <= len(nums); i++ {
		if freq[i] == 0 {
			ans = append(ans, i)
		}
	}

	return ans
}
