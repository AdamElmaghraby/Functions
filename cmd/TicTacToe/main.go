package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	playerX = "X"
	player0 = "0"
	empty   = " "
)

type Board [3][3]string

func initializedBoard() Board {
	var board Board
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			board[x][y] = empty
		}
	}
	return board
}

func printBoard(board Board) {
	fmt.Println("  0 1 2")
	fmt.Println(" -------")
	for x := 0; x < 3; x++ {
		fmt.Print(x)
		fmt.Print("|")
		for y := 0; y < 3; y++ {
			fmt.Printf(" %s", board[x][y])
		}
		fmt.Printf("|")
		fmt.Println("\n -------")
	}
}

func checkWin(board Board, player string) bool {
	for x := 0; x < 3; x++ {
		if (board[x][0] == player && board[x][1] == player && board[x][2] == player) ||
			(board[0][x] == player && board[1][x] == player && board[2][x] == player) {
			return true
		}
	}

	if (board[0][0] == player && board[1][1] == player && board[2][2] == player) ||
		(board[0][2] == player && board[1][1] == player && board[2][0] == player) {
		return true
	}

	return false
}

func isBoardFull(board Board) bool {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if board[x][y] == empty {
				return false
			}
		}
	}
	return true
}

func playerMove(board Board) (int, int) {
	var row, col int
	for {
		fmt.Print("Enter your move by row and column with a space seperated (for ex: '1 3'):\n")
		_, err := fmt.Scan(&row, &col)
		if err == nil && row >= 0 && row < 3 && col >= 0 && col < 3 && board[row][col] == empty {
			return row, col
		}
		fmt.Println("Invalid move. Try again!!\n (The cordinates must be in the bounds of the board and is empty)")
	}
}

func aiMove(board Board) (int, int) {
	rand.Seed(time.Now().UnixNano())
	var row, col int
	for {
		row = rand.Intn(3)
		col = rand.Intn(3)
		if board[row][col] == empty {
			return row, col
		}
	}
}

func main() {
	board := initializedBoard()
	currentPlayer := playerX

	for {
		printBoard(board)

		var row, col int
		if currentPlayer == playerX {
			row, col = playerMove(board)
		} else {
			fmt.Println("Computer's move")
			row, col = aiMove(board)
		}

		board[row][col] = currentPlayer

		if checkWin(board, currentPlayer) {
			printBoard(board)
			fmt.Printf("X wins!\n")
			break
		} else if isBoardFull(board) {
			printBoard(board)
			fmt.Println("It's a tie!")
			break
		}

		if currentPlayer == playerX {
			currentPlayer = player0
		} else {
			currentPlayer = playerX
		}
	}
}
