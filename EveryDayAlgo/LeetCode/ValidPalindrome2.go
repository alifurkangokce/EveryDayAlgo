package main

func validPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	left, right := 0, len(s)-1
	for left <= right && s[left] == s[right] {
		left++
		right--
	}

	if left >= right {
		return true
	}
	return isPalindromeHelper(string(s[:left])+string(s[left+1:])) ||
		isPalindromeHelper(string(s[:right])+string(s[right+1:]))
}

func isPalindromeHelper(s string) bool {
	if len(s) <= 1 {
		return true
	}

	return s[0] == s[len(s)-1] && isPalindromeHelper(s[1:len(s)-1])
}
