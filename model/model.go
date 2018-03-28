package model

import "errors"

// Model represents a game of tic tac toe
// Contains:
// a board to save state,
// a turn to know who can move,
// and names for the players
type Model struct {
	playerNames [2]string
	turn        bool
	board       [9]int
}

func NewGame(playerX, playerY string) Model {
	return Model{
		playerNames: [2]string{playerX, playerY},
		turn: true,
		board: [9]int{0,0,0,0,0,0,0,0,0}}
}

func (model *Model) MakeMove(x, y int) error {
	if x > 3 || x < 1 || y > 3 || y < 1 {
		return errors.New("move must be a coordinate (x, y) where x,y in [1,2,3]")
	}
	x = x - 1
	y = y - 1
	pos := x*3 + y

	if model.board[pos] == 1 || model.board[pos] == -1 {
		return errors.New("this position has already been taken")
	}

	if model.turn {
		model.board[pos] = 1
	} else {
		model.board[pos] = -1
	}
	model.turn = !model.turn
	return nil
}

func (model *Model) IsOver() bool {
	return false
}

func (model *Model) WhoWon() string {
	return model.playerNames[0]
}

func (model *Model) GetState() ([2]string, bool, [9]int) {
	return model.playerNames, model.turn, model.board
}

func (model *Model) Turn() string {
	if model.turn {
		return model.playerNames[0]
	}
	return model.playerNames[1]
}