package main

import (
	"fmt"

	"github.com/bdunbar22/TicTacToe/model"
	"github.com/bdunbar22/TicTacToe/view"
)

func main() {
	fmt.Print("Welcome to Tic-Tac-Toe")

	model := model.NewGame("Player 1", "Player 2")

	view.String(model)

	for model.IsOver() != true {
		// get input
		// update state
		// print state
	}

	//print winner.
 }