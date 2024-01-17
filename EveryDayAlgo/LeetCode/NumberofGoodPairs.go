package main

func numIdenticalPairs(nums []int) int {
	var res int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j <= len(nums)-1; j++ {
			if nums[i] == nums[j] {
				res++
			}
		}
	}
	return res

	// ans := 0
	// cnt := make(map[int]int)

	// for _, x := range A {
	//     ans += cnt[x]
	//     cnt[x]++
	// }

	// return ans
}
