package main

func isPalindromeLinked(head *ListNode) bool {
	// var res []int
	// dummy := head
	// for dummy != nil {
	// 	res = append(res, dummy.Val)
	// 	dummy = dummy.Next
	// }
	// j := len(res) - 1
	// for i := 0; i < len(res); i++ {
	// 	if res[i] != res[j] {
	// 		return false
	// 	}
	// 	j--
	// }
	// return true

	fast := head
	slow := head
	stack := []int{}

	for fast != nil && fast.Next != nil {
		stack = append(stack, slow.Val)
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	for slow != nil {
		if slow.Val != stack[len(stack)-1] {
			return false
		}
		slow = slow.Next
		stack = stack[:len(stack)-1]
	}
	return true
}
