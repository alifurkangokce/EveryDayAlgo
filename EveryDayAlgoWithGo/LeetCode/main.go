package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//fmt.Print(TwoSum([]int{3, 2, 4}, 6))
	// fmt.Print(isPalindrome(121))
	//fmt.Print(romanToInt("MCMXCIV"))
	//fmt.Print(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	//fmt.Print(isValid("(([]))"))
	//fmt.Print(mergeTwoLists(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}, &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}))
	//fmt.Print(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	//fmt.Print(strStr("hello", "ll"))
	//fmt.Print(searchInsert([]int{1, 3, 5, 6}, 2))
	//fmt.Print(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	//fmt.Print(lengthOfLastWord("   fly me   to   the moon  "))
	//fmt.Print(plusOne([]int{8}))
	//fmt.Print(addBinary("11", "1"))
	//fmt.Print(mySqrt(8))
	//fmt.Print(climbStairs(5))
	//merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	//fmt.Print(inorderTraversal(&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}))
	//fmt.Print(isSameTree(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}, &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}))
	//fmt.Print(isSymmetric(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}))
	//fmt.Print(maxDepth(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}))
	//fmt.Print(sortedArrayToBST([]int{-10, -3, 0, 5, 9}))
	//fmt.Print(isBalanced(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}))
	//fmt.Print(minDepth(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}))
	//fmt.Print(hasPathSum(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 3}}}, 6))
	//fmt.Print(generate(5))
	//fmt.Print(getRow(3))
	//fmt.Print(maxProfit([]int{7, 6, 4, 3, 1}))
	//fmt.Print(isPalindromeString("A man, a plan, a canal: Panama"))
	//fmt.Print(singleNumber([]int{1, 2, 1}))
	//fmt.Print(hasCycle(&ListNode{Val: 1, Next: &ListNode{Val: 2}}))
	//fmt.Print(preorderTraversal(&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}))
	//fmt.Print(postorderTraversal(&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}))
	//fmt.Print(getIntersectionNode(&ListNode{Val: 4, Next: &ListNode{Val: 1, Next: &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}, &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}))
	//fmt.Print(convertToTitle(701))
	//fmt.Print(titleToNumber("AB"))
	//fmt.Print(reverseBits(43261596))
	//fmt.Print(hammingWeight(11))
	//fmt.Print(isHappy(19))
	//fmt.Print(removeElements(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: nil}}}, 6))
	//fmt.Print(isIsomorphic("paper", "title"))
	//fmt.Print(reverseList(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: nil}}}))
	//fmt.Print(containsDuplicate([]int{1, 2, 3, 2}))
	//fmt.Print(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	// x := MyStack{}
	// x.Push(1)
	// x.Push(2)
	// fmt.Println(x.Top())
	// fmt.Println(x.Pop())
	// fmt.Println(x.Empty())
	//fmt.Print(invertTree(&TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}}}))
	//fmt.Print(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	//fmt.Print(isPowerOfTwo(2))

	// x := MyQueue{}
	// x.Push(1)
	// x.Push(2)
	// x.Peek()
	// x.Pop()
	// x.Empty()
	//fmt.Print(isPalindromeLinked(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}}}))
	//fmt.Print(isAnagram("rat", "car"))
	//fmt.Print(binaryTreePaths(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: nil, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}))
	fmt.Print(addDigits(38))
}
