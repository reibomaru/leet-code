package leetcode

import (
	"slices"
)

/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */

// @lc code=start


func wordBreak(s string, wordDict []string) bool {
	memo := make(map[int]bool)
	return searchDict(s, len(s)-1, &wordDict, memo)
}

func searchDict(s string, end int, dict *[]string, memo map[int]bool) bool {
	if end < 0 {
		return true
	}
	if memo, ok := memo[end]; ok {
		return memo
	}
	index := end
	for index >= 0 {
		if result := slices.Contains(*dict, s[index:end+1]); result {
			resp := searchDict(s, index-1, dict, memo)
			memo[index] = resp
			if resp {
				return true
			}
		}
		index--
	}
	return false
}
// @lc code=end

