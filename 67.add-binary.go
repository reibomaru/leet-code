package leetcode

import "strconv"

/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */

// @lc code=start
func addBinary(a string, b string) (result string) {
    aIndex := len(a) - 1
	bIndex := len(b) - 1

	carry := 0
	for aIndex >= 0 && bIndex >= 0 {
		var res int
		res, carry = culcBinary(a[aIndex], b[bIndex], carry)
		result = strconv.Itoa(res) + result
		aIndex--
		bIndex--
	}

	for aIndex >= 0 {
		var res int
		res, carry = culcBinary(a[aIndex], 0, carry)
		result = strconv.Itoa(res) + result
		aIndex--
	}

	for bIndex >= 0 {
		var res int
		res, carry = culcBinary(0,  b[bIndex], carry)
		result = strconv.Itoa(res) + result
		bIndex--
	}

	if carry == 1 {
		result = "1" + result
	}

	return result
}

func culcBinary(s1 byte, s2 byte, carry int) (int, int) {
	v1, _ := strconv.Atoi(string(s1))
	v2, _ := strconv.Atoi(string(s2))
	t := v1 + v2 + carry
	if t == 0 {
		return 0, 0
	}
	if t == 1 {
		return 1, 0
	}
	if t == 2 {
		return 0, 1
	}
	// t == 3
	return 1, 1
}
// @lc code=end

