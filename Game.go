package main

import "fmt"

//Game is a struct that represents the status of the game
type Game struct {
	board [8][8]int
	score map[int]int
}

var game = &Game{
	[8][8]int{},
	map[int]int{},
}

//SetupGame sets the board and score up for a new game
func SetupGame() {
	game.board[3][3] = 1
	game.board[4][4] = 1
	game.board[3][4] = 2
	game.board[4][3] = 2
	game.score[1] = 2
	game.score[2] = 2
	PrintGame()
}

//PrintGame prints out the board and score of the game
func PrintGame() {
	for _, row := range game.board {
		fmt.Println(row)
	}
	fmt.Printf("Player1: %v - Player2: %v\n", game.score[1], game.score[2])
}
