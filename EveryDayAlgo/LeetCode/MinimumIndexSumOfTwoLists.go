package main

func findRestaurant(list1 []string, list2 []string) []string {
	min := -1
	var res []string
	for i := 0; i <= len(list1)-1; i++ {
		for j := 0; j <= len(list2)-1; j++ {
			if list1[i] == list2[j] {
				if min == -1 {
					min = i + j
					res = append(res, list1[i])
				} else {
					if i+j == min {
						res = append(res, list1[i])
					} else if i+j < min {
						min = i + j
						res = []string{}
						res = append(res, list1[i])
					}
				}

			}
		}
	}
	return res

}
