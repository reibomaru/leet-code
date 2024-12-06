package leetcode

import "strings"

/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

// @lc code=start
func isPalindrome(s string) bool {
	str := strip(s)
    left := 0
	right := len(str) - 1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func strip(s string) string {
    var result strings.Builder
    for i := 0; i < len(s); i++ {
        b := s[i]
        if ('a' <= b && b <= 'z') ||
            ('A' <= b && b <= 'Z') || ('0' <= b && b <= '9') {
            result.WriteByte(b)
        }
    }
    return strings.ToLower(result.String())
}
// @lc code=end

