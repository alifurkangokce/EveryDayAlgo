package main

func lemonadeChange(bills []int) bool {
	if bills[0] != 5 {
		return false
	}
	var five, ten []int
	for i := 0; i <= len(bills)-1; i++ {
		if bills[i] == 5 {
			five = append(five, bills[i])
		} else if bills[i] == 10 {
			if len(five) == 0 {
				return false
			}
			five = five[1:]
			ten = append(ten, bills[i])
		} else {
			if len(five) == 0 || (len(ten) == 0 && len(five) < 3) {
				return false
			}
			if len(ten) > 0 {
				ten = ten[1:]
				five = five[1:]
			} else {
				five = five[3:]
			}

		}
	}
	return true
}
