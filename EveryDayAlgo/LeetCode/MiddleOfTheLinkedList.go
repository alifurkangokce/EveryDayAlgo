package main

func middleNode(head *ListNode) *ListNode {
	var arr []int
	var res *ListNode = head
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	for i := 0; i < len(arr)/2; i++ {
		res = res.Next
	}
	return res
}

// func middleNode(head *ListNode) *ListNode {
//     fp := head
//     sp := head
//     for fp != nil && fp.Next != nil {
//         sp = sp.Next
//         fp = fp.Next.Next
//     }
//     return sp
// }
