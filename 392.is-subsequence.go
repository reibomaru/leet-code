package leetcode

/*
 * @lc app=leetcode id=392 lang=golang
 *
 * [392] Is Subsequence
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}
    sIndex := 0
	for i := range len(t) {
		if t[i] == s[sIndex] {
			sIndex++
			if sIndex > len(s) - 1 {
				return true
			}
		}
	}
	return false
}
// @lc code=end

