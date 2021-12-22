package main

import "fmt"

func main() {
	const COL = 3
	const ROW = 3
	var sampleBoard [ROW][COL]string = [ROW][COL]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	var board [ROW][COL]string = [ROW][COL]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}
	var number int
	var str string
	var row int
	var col int
	isFirst := true
	isGameEnd := false

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

		switch number {
		case 1:
			row = 0
			col = 0
		case 2:
			row = 0
			col = 1
		case 3:
			row = 0
			col = 2
		case 4:
			row = 1
			col = 0
		case 5:
			row = 1
			col = 1
		case 6:
			row = 1
			col = 2
		case 7:
			row = 2
			col = 0
		case 8:
			row = 2
			col = 1
		case 9:
			row = 2
			col = 2
		}

		switch isFirst {
		case true:
			str = "○"
		case false:
			str = "×"
		}

		if board[row][col] == "-" {
			board[row][col] = str
			fmt.Printf("\n")
			for i := 0; i < len(board); i++ {
				fmt.Printf("           ")
				fmt.Println(board[i])
			}
			fmt.Printf("\n")

			for i := 0; i < len(board); i++ {
				// row
				if board[i][0] != "-" &&
					board[i][1] != "-" &&
					board[i][2] != "-" &&
					board[i][0] == board[i][1] &&
					board[i][1] == board[i][2] {
					isGameEnd = true
				}
				// col
				if board[0][i] != "-" &&
					board[1][i] != "-" &&
					board[2][i] != "-" &&
					board[0][i] == board[1][i] &&
					board[1][i] == board[2][i] {
					isGameEnd = true
				}
			}
			// cross
			if board[0][0] != "-" &&
				board[1][1] != "-" &&
				board[2][2] != "-" &&
				board[0][0] == board[1][1] &&
				board[1][1] == board[2][2] {
				isGameEnd = true
			}
			if board[0][2] != "-" &&
				board[1][1] != "-" &&
				board[2][0] != "-" &&
				board[0][2] == board[1][1] &&
				board[1][1] == board[2][0] {
				isGameEnd = true
			}
			if isGameEnd {
				fmt.Printf("Game end\n")
				return
			}

			isFirst = !isFirst
		} else {
			fmt.Printf("\n")
			fmt.Printf("-------------------------------------\n")
			fmt.Printf("      already input the area!!!      \n")
			fmt.Printf("-------------------------------------\n")
			fmt.Printf("\n")
			for i := 0; i < len(board); i++ {
				fmt.Printf("           ")
				fmt.Println(board[i])
			}
		}
	}
}
