package model

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
	return Model{playerNames: [2]string{playerX, playerY}}
}

func (model *Model) MakeMove(x, y int) {

	pos := x*3 + y
	if model.turn {
		model.board[pos] = 1
	} else {
		model.board[pos] = -1
	}
	model.turn = !model.turn
}

func (model Model) IsOver() bool {
	return false
}

func (model Model) WhoWon() string {
	return model.playerNames[0]
}

func (model *Model) GetState() ([2]string, bool, [9]int) {
	return model.playerNames, model.turn, model.board
}