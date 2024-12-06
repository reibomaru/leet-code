package leetcode

/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
    min, left, sum := 0, 0, 0

    for right := range nums {
        sum += nums[right]
        for sum >= target{
            size := right - left + 1
			if min == 0 || size < min{
				min = size
			}
            sum -= nums[left]
            left++
        }
    }
    return min
}
// @lc code=end

