//chessboard with arrays
package main

import (
	"fmt"
)

//func chessboard() {

//}

func main() {
	var board [8][8]rune

	//Black pieces
	board[0][0] = 'r'
	board[0][1] = 'n'
	board[0][2] = 'b'
	board[0][3] = 'q'
	board[0][4] = 'k'
	board[0][5] = 'b'
	board[0][6] = 'n'
	board[0][7] = 'r'

	//White pieces
	board[7][0] = 'R'
	board[7][1] = 'N'
	board[7][2] = 'B'
	board[7][3] = 'Q'
	board[7][4] = 'K'
	board[7][5] = 'B'
	board[7][6] = 'N'
	board[7][7] = 'R'

	for _, chessboard := range board {
		for _, row := range chessboard {
			if row == 0 {
				fmt.Print("|_| ")
			} else {
				fmt.Printf("|%c| ", row)
			}
		}
		fmt.Print("\n")
	}

}
