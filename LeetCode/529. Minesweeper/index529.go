package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strconv"
)

func print(matrix [][]byte) {
	for _, val := range matrix {
		for _, val2 := range val {
			fmt.Printf(string(val2) + " ")
		}
		fmt.Println()
	}
}

func main() {
	// print(updateBoard([][]byte{
	// 	{'E', 'E', 'E', 'E', 'E'},
	// 	{'E', 'E', 'M', 'E', 'E'},
	// 	{'E', 'E', 'E', 'E', 'E'},
	// 	{'E', 'E', 'E', 'E', 'E'},
	// }, []int{3, 0}))

	// print(updateBoard([][]byte{
	// 	{'E', 'E', 'E', 'E', 'E', 'E', 'E', 'E'},
	// 	{'E', 'E', 'E', 'E', 'E', 'E', 'E', 'M'},
	// 	{'E', 'E', 'M', 'E', 'E', 'E', 'E', 'E'},
	// 	{'M', 'E', 'E', 'E', 'E', 'E', 'E', 'E'},
	// 	{'E', 'E', 'E', 'E', 'E', 'E', 'E', 'E'},
	// 	{'E', 'E', 'E', 'E', 'E', 'E', 'E', 'E'},
	// 	{'E', 'E', 'E', 'E', 'E', 'E', 'E', 'E'},
	// 	{'E', 'E', 'M', 'M', 'E', 'E', 'E', 'E'},
	// }, []int{0, 0}))
	// [
	// ["B","B","B","B","B","B","1","E"],
	// ["B","1","1","1","B","B","1","M"],
	// ["1","2","M","1","B","B","1","1"],
	// ["M","2","1","1","B","B","B","B"],
	// ["1","1","B","B","B","B","B","B"],
	// ["B","B","B","B","B","B","B","B"],
	// ["B","1","2","2","1","B","B","B"],
	// ["B","1","M","M","1","B","B","B"]
	// ]
	print(updateBoard([][]byte{
		{'B', 'B', 'B', 'B', '1', 'M', 'M', 'E'},
		{'B', 'B', 'B', 'B', '1', '4', 'M', 'E'},
		{'B', 'B', 'B', 'B', 'B', '3', 'M', 'E'},
		{'B', 'B', 'B', 'B', 'B', '2', 'M', 'E'},
		{'1', '2', '2', '1', 'B', '1', '1', '1'},
		{'E', 'M', 'M', '1', 'B', 'B', 'B', 'B'},
		{'E', 'E', 'E', '2', '2', '2', '2', '1'},
		{'E', 'E', 'E', 'E', 'M', 'M', 'E', 'M'},
	}, []int{7, 2}))
	// [
	// ["B","B","B","B","1","M","M","E"],
	// ["B","B","B","B","1","4","M","E"],
	// ["B","B","B","B","B","3","M","E"],
	// ["B","B","B","B","B","2","M","E"],
	// ["1","2","2","1","B","1","1","1"],
	// ["E","M","M","1","B","B","B","B"],
	// ["1","2","2","2","2","2","2","1"],
	// ["B","B","B","1","M","M","E","M"]
	// ]
}

func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func getMineNum(board [][]byte, i, j int) byte {
	ct := 0
	lenx := len(board)
	for x := -1; x <= 1; x++ {
		leny := len(board[i])
		for y := -1; y <= 1; y++ {
			if i+x >= 0 && i+x < lenx && y+j >= 0 && y+j < leny && (x != 0 || y != 0) {
				if board[i+x][j+y] == 'M' {
					ct++
				}
			}
		}
	}
	return strconv.Itoa(ct)[0]
}

func updateBoard(board [][]byte, click []int) [][]byte {
	lenx := len(board)
	var preHandleBoard [][]byte
	if err := deepCopy(&preHandleBoard, &board); err != nil {
		panic(err.Error())
	}
	record := make([][]int, len(board))

	print(preHandleBoard)
	println("")
	for i, vali := range preHandleBoard {
		leny := len(preHandleBoard[i])
		record[i] = make([]int, leny)
		for j, valj := range vali {
			if valj == 'M' {
				// 将四周的填上数字
				for x := -1; x <= 1; x++ {
					for y := -1; y <= 1; y++ {
						if i+x >= 0 && i+x < lenx && y+j >= 0 && y+j < leny && (x != 0 || y != 0) {
							if preHandleBoard[i+x][j+y] == 'E' {
								// 需要 填充
								preHandleBoard[i+x][j+y] = getMineNum(board, i+x, j+y)
							}
						}
					}
				}
			}
		}
	}

	print(preHandleBoard)
	// 然后对点击的方块处理
	block := preHandleBoard[click[0]][click[1]]

	switch block {
	case 'E':
		// 深搜替换成B
		deepReplace(&record, &board, &preHandleBoard, click[0], click[1])
		break
	case 'M':
		// 替换成X
		board[click[0]][click[1]] = 'X'
		break
	default: // 数字
		board[click[0]][click[1]] = preHandleBoard[click[0]][click[1]]
	}
	return board
}

func deepReplace(record *[][]int, board, preHandleBoard *[][]byte, i, j int) {
	if i < 0 || i >= len(*board) || j < 0 || j >= len((*board)[0]) {
		return
	}
	if (*record)[i][j] == 1 {
		return
	}
	(*record)[i][j] = 1

	if (*preHandleBoard)[i][j] == 'E' {
		(*board)[i][j] = 'B'
	} else { // 数字
		(*board)[i][j] = (*preHandleBoard)[i][j]
		return
	}
	// fmt.Println(i, j)
	// print(*board)
	// fmt.Println("下")
	deepReplace(record, board, preHandleBoard, i+1, j)
	// fmt.Println("上")
	deepReplace(record, board, preHandleBoard, i-1, j)
	// fmt.Println("右")
	deepReplace(record, board, preHandleBoard, i, j+1)
	// fmt.Println("左")
	deepReplace(record, board, preHandleBoard, i, j-1)
	// 4个对角线
	deepReplace(record, board, preHandleBoard, i-1, j-1)
	deepReplace(record, board, preHandleBoard, i+1, j-1)
	deepReplace(record, board, preHandleBoard, i-1, j+1)
	deepReplace(record, board, preHandleBoard, i+1, j+1)
}
