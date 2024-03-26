package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var humanPlayer, computerPlayer rune

func getChoosenPlayer() rune {

	var player rune
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Choose Player (x or o): ")

	for {
		if scanner.Scan() {
			input := strings.TrimSpace(scanner.Text())

			if len(input) == 1 && (input == "x" || input == "o") {
				player = rune(input[0])
				break
			} else {
				fmt.Println("Invalid choice. Please choose 'x' or 'o'.")
			}
		}
	}

	fmt.Println("You play as:", string(player))
	return player

}

func getChoosenLocation() [2]int {

	var location [2]int
	var err error

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter x-coordinate: ")

	for {
		row, _ := reader.ReadString('\n')
		row = row[:len(row)-1]
		location[0], err = strconv.Atoi(row)
		if err != nil {
			fmt.Println("Enter a valid number")
		} else if location[0] > 3 {
			fmt.Println("x-coordinate cannot be greater than 3. Please enter a valid number")
		} else if location[0] < 1 {
			fmt.Println("x-coordinate cannot be less than 1. Please enter a valid number")
		} else {
			break
		}
	}

	fmt.Print("Enter y-coordinate: ")

	for {
		col, _ := reader.ReadString('\n')
		col = col[:len(col)-1]
		location[1], err = strconv.Atoi(col)
		if err != nil {
			fmt.Println("Enter a valid number")
		} else if location[1] > 3 {
			fmt.Println("y-coordinate cannot be greater than 3. Please enter a valid number")
		} else if location[0] < 1 {
			fmt.Println("x-coordinate cannot be less than 1. Please enter a valid number")
		} else {
			break
		}
	}

	return location

}

// func isLocationOccupied(board [][]rune, location [2]int) bool {
// 	row, col := location[0], location[1]
// 	return row >= 0 && row < len(board) && col >= 0 && col < len(board[0]) && board[row][col] != '_'
// }

func fillBoard(board [][]rune, location [2]int, player rune, isHumanInput bool) [][]rune {

	var row, col int

	if isHumanInput {
		row = location[0] - 1
		col = location[1] - 1

	} else {
		row = location[0]
		col = location[1]
	}

	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) {
		fmt.Println("Invalid coordinates. Please enter values between 1 and 3.")
		return board
	}

	if board[row][col] != '_' {
		fmt.Printf("Location (%d, %d) is already occupied\n", row+1, col+1)
		return board
	}

	board[row][col] = player
	return board

}

func printBoard(board [][]rune) {

	for _, row := range board {
		for _, cell := range row {
			fmt.Printf("%c ", cell)
		}
		fmt.Println()
	}
}

func gameStatus(board [][]rune) (rune, bool) {

	for _, row := range board {
		if areEqual(row) {
			return row[0], true
		}
	}

	for col := range board[0] {

		column := make([]rune, len(board))

		for row := range board {
			column[row] = board[row][col]

			if areEqual(column) {

				return column[0], true
			}
		}

	}

	if len(board) == len(board[0]) {

		if board[0][0] != '_' {
			if board[0][0] == board[1][1] && board[1][1] == board[2][2] {

				return board[0][0], true
			}
		}

		if board[0][2] != '_' {
			if board[0][2] == board[1][1] && board[1][1] == board[2][0] {

				return board[0][2], true
			}
		}

		noMoves := isMovesLeft(board)

		if noMoves == false {
			return ' ', true
		}

	}

	return '_', false
}

func areEqual(slice []rune) bool {

	if len(slice) == 0 {
		return false
	}

	standardValue := slice[0]

	for _, element := range slice[1:] {
		if element != standardValue || element == '_' {
			return false
		}
	}

	return true
}

func evaluate(board [][]rune) int {

	winner, gameOver := gameStatus(board)

	if gameOver {
		if winner == computerPlayer {
			return +10
		} else if winner == humanPlayer {
			return -10
		}
	}

	return 0

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isMovesLeft(board [][]rune) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == '_' {
				return true
			}
		}
	}
	return false
}

func minimax(board [][]rune, depth int, maximizingPlayer bool) int {

	score := evaluate(board)

	if score == 10 {
		return score
	}
	if score == -10 {
		return score
	}
	if !isMovesLeft(board) {
		return 0
	}

	if maximizingPlayer {
		bestValue := math.MinInt32
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == '_' {

					board[i][j] = computerPlayer
					bestValue = max(bestValue, minimax(board, depth+1, false))
					board[i][j] = '_'
				}
			}
		}
		return bestValue
	} else {
		bestValue := math.MaxInt32
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == '_' {

					board[i][j] = humanPlayer
					bestValue = min(bestValue, minimax(board, depth+1, true))
					board[i][j] = '_'
				}
			}
		}
		return bestValue
	}
}

func findBestMove(board [][]rune) [2]int {

	var bestMove [2]int
	best := math.MinInt32

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == '_' {

				board[i][j] = computerPlayer

				valueForMove := minimax(board, 0, false)

				board[i][j] = '_'

				if valueForMove > best {
					bestMove = [2]int{i, j}
					best = valueForMove
				}

			}
		}
	}

	return bestMove

}

func main() {

	board := [][]rune{
		{'_', '_', '_'},
		{'_', '_', '_'},
		{'_', '_', '_'},
	}

	var whoseTurne bool

	player := getChoosenPlayer()

	if player == 'x' {
		humanPlayer = 'x'
		computerPlayer = 'o'
		whoseTurne = true
	} else {
		humanPlayer = 'o'
		computerPlayer = 'x'
		whoseTurne = false
	}

	for {

		winner, gameOver := gameStatus(board)

		if gameOver {
			if winner == ' ' {
				fmt.Println("No winner")
			} else {
				fmt.Println("Winner:", string(winner))
			}
			break
		}

		if whoseTurne { // Human player's turn
			location := getChoosenLocation()
			board = fillBoard(board, location, humanPlayer, true)
			printBoard(board)

		} else { // Computer's turn

			bestMove := findBestMove(board)
			fmt.Println(bestMove)
			board = fillBoard(board, bestMove, computerPlayer, false)
			printBoard(board)

		}

		whoseTurne = !whoseTurne

	}

}
