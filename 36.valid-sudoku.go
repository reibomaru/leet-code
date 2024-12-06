package leetcode

import "fmt"

/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
    verticalMap := make(map[int]map[byte]int)
	horizontalMap := make(map[int]map[byte]int)
	areaMap := make(map[string]map[byte]int)

	var val byte
	for i := range 9 {
		for j := range 9{
			if val = board[i][j]; val == '.' {
				continue
			}
			if _, ok := verticalMap[j][val]; ok {
				return false
			}else{
				if _, ok := verticalMap[j]; !ok {
					verticalMap[j] = make(map[byte]int) // ここで初期化
				}
				verticalMap[j][val] = 1
			}
			if _, ok := horizontalMap[i][val]; ok {
				return false
			}else{
				if _, ok := horizontalMap[i]; !ok {
					horizontalMap[i] = make(map[byte]int) // ここで初期化
				}
				horizontalMap[i][val] = 1
			}
			area := area(i,j)
			if _, ok := areaMap[area][val]; ok {
				return false
			}else{
				if _, ok := areaMap[area]; !ok {
					areaMap[area] = make(map[byte]int) // ここで初期化
				}
				areaMap[area][val] = 1
			}
		}
	}
	return true

}

func area(i, j int) string {
	return fmt.Sprintf("%d-%d", i / 3, j / 3)
}
// @lc code=end

