package main

import (
	"fmt"
)
func main(){
	fmt.Println(isValidSudoku([][]byte{
		{'5','3','.','.','7','.','.','.','.'},
		{'6','.','.','1','9','5','.','.','.'},
		{'.','9','8','.','.','.','.','6','.'},
		{'8','.','.','.','6','.','.','.','3'},
		{'4','.','.','8','.','3','.','.','1'},
		{'7','.','.','.','2','.','.','.','6'},
		{'.','6','.','.','.','.','2','8','.'},
		{'.','.','.','4','1','9','.','.','5'},
		{'.','.','.','.','8','.','.','7','9'},
	}))
	fmt.Println(isValidSudoku([][]byte{
	{'8','3','.','.','7','.','.','.','.'},
	{'6','.','.','1','9','5','.','.','.'},
	{'.','9','8','.','.','.','.','6','.'},
	{'8','.','.','.','6','.','.','.','3'},
	{'4','.','.','8','.','3','.','.','1'},
	{'7','.','.','.','2','.','.','.','6'},
	{'.','6','.','.','.','.','2','8','.'},
	{'.','.','.','4','1','9','.','.','5'},
	{'.','.','.','.','8','.','.','7','9'},
	}))
	fmt.Println(isValidSudoku([][]byte{
		{'.','.','.','8','.','.','.','.','.'},
		{'.','.','.','.','.','.','.','.','.'},
		{'.','6','.','.','.','.','3','.','.'},
		{'7','.','.','9','6','4','1','.','.'},
		{'6','.','9','.','.','.','.','.','.'},
		{'.','.','.','.','.','.','.','5','.'},
		{'.','.','9','.','.','.','.','.','.'},
		{'.','.','.','.','.','.','.','.','5'},
		{'.','.','1','.','.','.','.','2','.'},
	}))
}

func isValidSudoku(board [][]byte) bool {
	var i,j,k,t int
	for i=0;i<9;i++{
		for j=0;j<9;j++{
			if board[i][j] != '.'{
				//行
				ct := 0
				for k=0;k<9;k++{
					if board[i][k] == board[i][j]{
						ct++
					}
				}
				if ct >= 2 {
					return false 
				}
				//列
				ct = 0
				for k=0;k<9;k++{
					if board[k][j] == board[i][j]{
						ct++
					}
				}
				if ct >= 2 {
					return false 
				}
				//区块
				ct =0
				startX := i-i%3
				startY := j-j%3
				for k=0;k<3;k++{
					for t=0;t<3;t++{
						if board[k+startX][t+startY] == board[i][j]{
							ct ++ 
						}
					}
				}
				if ct >= 2 {
					return false 
				}
			}
		}
	}

	return true	
}
// Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition.

// A partially filled sudoku which is valid.

// The Sudoku board could be partially filled, where empty cells are filled with the character '.'.

// Example 1:

// Input:
// [
//   ["5","3",".",".","7",".",".",".","."],
//   ["6",".",".","1","9","5",".",".","."],
//   [".","9","8",".",".",".",".","6","."],
//   ["8",".",".",".","6",".",".",".","3"],
//   ["4",".",".","8",".","3",".",".","1"],
//   ["7",".",".",".","2",".",".",".","6"],
//   [".","6",".",".",".",".","2","8","."],
//   [".",".",".","4","1","9",".",".","5"],
//   [".",".",".",".","8",".",".","7","9"]
// ]
// Output: true
// Example 2:

// Input:
// [
//   ["8","3",".",".","7",".",".",".","."],
//   ["6",".",".","1","9","5",".",".","."],
//   [".","9","8",".",".",".",".","6","."],
//   ["8",".",".",".","6",".",".",".","3"],
//   ["4",".",".","8",".","3",".",".","1"],
//   ["7",".",".",".","2",".",".",".","6"],
//   [".","6",".",".",".",".","2","8","."],
//   [".",".",".","4","1","9",".",".","5"],
//   [".",".",".",".","8",".",".","7","9"]
// ]
// Output: false
// Explanation: Same as Example 1, except with the 5 in the top left corner being 
//     modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.
// Note:

// A Sudoku board (partially filled) could be valid but is not necessarily solvable.
// Only the filled cells need to be validated according to the mentioned rules.
// The given board contain only digits 1-9 and the character '.'.
// The given board size is always 9x9.