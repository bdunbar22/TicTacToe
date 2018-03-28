package controller

// Controller is used to get input from human players and/or computer players

import (
	"github.com/bdunbar22/TicTacToe/model"
	"bufio"
	"os"
	"fmt"
	"strconv"
)

func GetMove(model model.Model) (x, y int) {
	reader := bufio.NewReader(os.Stdin)
	loopCondition := true
	for loopCondition {
		fmt.Printf("%v - Enter move:", model.Turn())
		text, _ := reader.ReadString('\n')
		if len(text) < 3 || string(text[1]) != " " {
			fmt.Print("Invalid input please input in the format: '# #'")
		} else {
			loopCondition = false
			x, _ = strconv.Atoi(string(text[0]))
			y, _ = strconv.Atoi(string(text[2]))
		}
	}
	return
}