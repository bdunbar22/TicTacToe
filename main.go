package main

import (
	"fmt"

	"github.com/bdunbar22/TicTacToe/model"
	"github.com/bdunbar22/TicTacToe/view"
	"github.com/bdunbar22/TicTacToe/controller"
)

func main() {
	fmt.Println("Welcome to Tic-Tac-Toe")

	model := model.NewGame("Player 1", "Player 2")

	view.String(model)

	i := 0

	for model.IsOver() != true && i < 9 {
		loopCondition := true
		for loopCondition {
			x, y := controller.GetMove(model)
			message := model.MakeMove(x, y)
			if message != nil {
				fmt.Println(message.Error())
			} else {
				loopCondition = false
			}
		}
		fmt.Println(view.String(model))
		i ++
	}

	//print winner.
 }