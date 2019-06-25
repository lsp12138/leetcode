package main

import "fmt"

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board))
}
func isValidSudoku(board [][]byte) bool {
	// 初始都为false
	row, column, box := [9][9]bool{}, [9][9]bool{}, [9][9]bool{}
	for r := range board {
		for c := range board[r] {
			if board[r][c] == '.' {
				continue
			}
			// 数字的下标为数字的值减一
			digit := board[r][c] - '0' - 1
			// 9个3x3的区域的下标
			boxIdx := 3*(r/3) + c/3
			if row[r][digit] || column[c][digit] || box[boxIdx][digit] {
				return false
			}
			// true表示该数出现过
			row[r][digit] = true
			column[c][digit] = true
			box[boxIdx][digit] = true
		}
	}
	return true
}
