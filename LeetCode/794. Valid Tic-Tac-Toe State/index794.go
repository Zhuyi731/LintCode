package main

import "fmt"

func main() {
	fmt.Println(validTicTacToe([]string{"XO ", "   ", "X  "}))
	fmt.Println(validTicTacToe([]string{
		"OXX",
		"XOX",
		"OXO"}))
	fmt.Println(validTicTacToe([]string{"O  ", "   ", "   "}))
	fmt.Println(validTicTacToe([]string{"XOX", " X ", "   "}))
	fmt.Println(validTicTacToe([]string{"XXX", "   ", "OOO"}))
	fmt.Println(validTicTacToe([]string{"XOX", "O O", "XOX"}))
	fmt.Println(validTicTacToe([]string{"XOX", "OOO", "XOX"}))
	fmt.Println(validTicTacToe([]string{"XXX", "XOO", "OO "}))

}

func validTicTacToe(board []string) bool {
	// 规则
	xCt := 0
	oCt := 0
	fillCt := []int{0, 0}
	for _, str := range board {
		for _, r := range str {
			if r == 'X' {
				xCt++
			} else if r == 'O' {
				oCt++
			}
		}
	}

	if xCt-oCt < 0 || xCt-oCt > 1 {
		return false
	}

	for i := 0; i < 3; i++ {
		if board[i] == "XXX" {
			fillCt[0]++
		}
		if board[i] == "OOO" {
			fillCt[1]++
		}
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			if board[0][i] == 'X' {
				fillCt[0]++
			} else if board[0][i] == 'O' {
				fillCt[1]++
			}
		}
	}
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		if board[1][1] == 'X' {
			fillCt[0]++
		} else if board[1][1] == 'O' {
			fillCt[1]++
		}
	}
	if board[2][0] == board[1][1] && board[1][1] == board[0][2] {
		if board[1][1] == 'X' {
			fillCt[0]++
		} else if board[1][1] == 'O' {
			fillCt[1]++
		}
	}
	if fillCt[0] != 0 && fillCt[1] != 0 {
		return false
	}
	if fillCt[0] > 0 && xCt-oCt != 1 {
		return false
	}
	if fillCt[1] > 0 && xCt-oCt != 0 {
		return false
	}

	return true
}
