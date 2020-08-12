package main

import "fmt"

func main() {
	// testCase := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	testCase := [][]byte{
		{'O', 'X', 'X', 'O', 'X'},
		{'X', 'O', 'O', 'X', 'O'},
		{'X', 'O', 'X', 'O', 'X'},
		{'O', 'X', 'O', 'O', 'O'},
		{'X', 'X', 'O', 'X', 'O'},
	}
	// ["O","X","X","O","X"],
	// ["X","X","X","X","O"],
	// ["X","X","X","X","X"],
	// ["O","X","O","X","X"],
	// ["X","X","O","X","X"]
	solve(testCase)
}

func print(board [][]byte) {
	var lenI = len(board)
	var lenJ = len(board[0])
	for i := 0; i < lenI; i++ {
		for j := 0; j < lenJ; j++ {
			fmt.Print(string(board[i][j]))
		}
		fmt.Println()
	}
}

var checked [][]byte

var lenI, lenJ int

func solve(board [][]byte) {
	var i, j int
	lenI = len(board)
	if lenI == 0 {
		return
	}
	lenJ = len(board[0])

	for i = 0; i < lenI; i++ {
		if board[i][0] == 'O' {
			deepReplace(board, i, 0)
		}
		if board[i][lenJ-1] == 'O' {
			deepReplace(board, i, lenJ-1)
		}
	}

	for j = 0; j < lenJ; j++ {
		if board[0][j] == 'O' {
			deepReplace(board, 0, j)
		}
		if board[lenI-1][j] == 'O' {
			deepReplace(board, lenI-1, j)
		}
	}

	for i = 0; i < lenI; i++ {
		for j = 0; j < lenJ; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '-' {
				board[i][j] = 'O'
			}
		}
	}

}

func deepReplace(board [][]byte, i, j int) {
	board[i][j] = '-'

	if i-1 >= 0 && board[i-1][j] == 'O' { // 上
		deepReplace(board, i-1, j)
	}

	if i+1 < lenI && board[i+1][j] == 'O' { // 下
		deepReplace(board, i+1, j)
	}

	if j-1 >= 0 && board[i][j-1] == 'O' { // 左
		deepReplace(board, i, j-1)
	}

	if j+1 < lenJ && board[i][j+1] == 'O' { // 右
		deepReplace(board, i, j+1)
	}
}
