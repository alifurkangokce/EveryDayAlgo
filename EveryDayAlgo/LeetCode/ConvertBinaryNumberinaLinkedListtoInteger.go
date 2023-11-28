package main

func getDecimalValue(head *ListNode) (decimal int) {
	for head != nil {
		decimal = (decimal * 2) + head.Val
		head = head.Next
	}
	return decimal
}
