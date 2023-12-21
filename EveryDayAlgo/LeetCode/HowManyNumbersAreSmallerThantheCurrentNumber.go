package main

func smallerNumbersThanCurrent(nums []int) []int {
	var res []int
	for i := 0; i <= len(nums)-1; i++ {
		var count int
		for j := 0; j <= len(nums)-1; j++ {
			if nums[i] > nums[j] {
				count++
			}
		}
		res = append(res, count)
	}
	return res
}

// int[] count = new int[101];
// int[] res = new int[nums.length];

// for (int i =0; i < nums.length; i++) {
// 	count[nums[i]]++;
// }

// for (int i = 1 ; i <= 100; i++) {
// 	count[i] += count[i-1];
// }

// for (int i = 0; i < nums.length; i++) {
// 	if (nums[i] == 0)
// 		res[i] = 0;
// 	else
// 		res[i] = count[nums[i] - 1];
// }

// return res;
