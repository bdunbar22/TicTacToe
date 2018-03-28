package view

// This view will return the game state in a readable string format.

import (
	"github.com/bdunbar22/TicTacToe/model"
	"bytes"
)

// String will return the model state formatted as a string
// Spaces will be X, O, or a blank space
func String(model model.Model) string {
	_, _, board := model.GetState()

	var b bytes.Buffer
	b.WriteString("======================\n")

	b.WriteString("-------\n")
	for i := 0; i < 3; i++ {

		for j := 0; j < 3; j++ {
			b.WriteString("|")
			switch {
			case board[3*i + j] < 0:
				b.WriteString("X")
			case board[3*i + j] == 0:
				b.WriteString(" ")
			case board[3*i + j] > 0:
				b.WriteString("O")
			}
		}
		b.WriteString("|\n")
		b.WriteString("-------\n")
	}

	return b.String()
}
