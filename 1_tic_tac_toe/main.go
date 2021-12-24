package main

import "fmt"

const COL = 3
const ROW = 3

type Game struct {
	isFirstPlayer bool
	gameBoard     [][]string
}

func isGameCompleted(board [][]string) bool {
	isGameCompleted := false
	for i := 0; i < len(board); i++ {
		// row
		if board[i][0] != "-" &&
			board[i][1] != "-" &&
			board[i][2] != "-" &&
			board[i][0] == board[i][1] &&
			board[i][1] == board[i][2] {
			isGameCompleted = true
		}
		// col
		if board[0][i] != "-" &&
			board[1][i] != "-" &&
			board[2][i] != "-" &&
			board[0][i] == board[1][i] &&
			board[1][i] == board[2][i] {
			isGameCompleted = true
		}
	}
	// cross
	if board[0][0] != "-" &&
		board[1][1] != "-" &&
		board[2][2] != "-" &&
		board[0][0] == board[1][1] &&
		board[1][1] == board[2][2] {
		isGameCompleted = true
	}
	if board[0][2] != "-" &&
		board[1][1] != "-" &&
		board[2][0] != "-" &&
		board[0][2] == board[1][1] &&
		board[1][1] == board[2][0] {
		isGameCompleted = true
	}
	return isGameCompleted
}

func main() {
	var sampleBoard [][]string = [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	var number int
	var str string
	var row int
	var col int
	gameInfo := new(Game)
	gameInfo.isFirstPlayer = true
	gameInfo.gameBoard = [][]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}

	fmt.Printf("--------------------------------------\n")
	for i := 0; ; i++ {
		fmt.Printf("\n")
		for i := 0; i < len(sampleBoard); i++ {
			fmt.Printf("           ")
			fmt.Println(sampleBoard[i])
		}
		fmt.Printf("\ninput the number: ")
		fmt.Scan(&number)
		if number < 1 || number > 9 {
			break
		}

		row = (number - 1) / 3
		col = (number - 1) % 3
		switch gameInfo.isFirstPlayer {
		case true:
			str = "○"
		case false:
			str = "×"
		}

		// Already entered
		if gameInfo.gameBoard[row][col] != "-" {
			fmt.Printf("\n")
			fmt.Printf("-------------------------------------\n")
			fmt.Printf("           Already entered           \n")
			fmt.Printf("-------------------------------------\n")
			fmt.Printf("\n")
			for i := 0; i < len(gameInfo.gameBoard); i++ {
				fmt.Printf("           ")
				fmt.Println(gameInfo.gameBoard[i])
			}
			continue
		}

		gameInfo.gameBoard[row][col] = str
		fmt.Printf("\n")
		for i := 0; i < len(gameInfo.gameBoard); i++ {
			fmt.Printf("           ")
			fmt.Println(gameInfo.gameBoard[i])
		}
		fmt.Printf("\n")
		if isGameCompleted(gameInfo.gameBoard) {
			fmt.Printf("Game end\n")
			return
		}

		gameInfo.isFirstPlayer = !gameInfo.isFirstPlayer

	}
}
