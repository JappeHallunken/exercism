package stateoftictactoe

import "fmt"

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

func StateOfTicTacToe(board []string) (State, error) {
	// count X and O
	var x, o int
	for _, v := range board {
		for _, r := range v {
			switch r {
			case 'X':
				x++
			case 'O':
				o++
			case ' ':
				continue
			default: //r is other then X, O or space
				return "", fmt.Errorf("invalid board")
			}
		}
	}
	//check correct order
	if o > x || x-o > 1 {
		return "", fmt.Errorf("wrong order")
	}
	//check win
	lines := [8][3][2]int{
		{{0, 0}, {0, 1}, {0, 2}}, {{1, 0}, {1, 1}, {1, 2}}, {{2, 0}, {2, 1}, {2, 2}}, // horizontal
		{{0, 0}, {1, 0}, {2, 0}}, {{0, 1}, {1, 1}, {2, 1}}, {{0, 2}, {1, 2}, {2, 2}}, // vertical
		{{0, 0}, {1, 1}, {2, 2}}, {{0, 2}, {1, 1}, {2, 0}}, // diagonal
	}
	xWins, oWins := 0, 0
	for _, l := range lines {
		a, b, c := board[l[0][0]][l[0][1]], board[l[1][0]][l[1][1]], board[l[2][0]][l[2][1]]
		if a == b && b == c {
			switch a {
			case 'X':
				xWins++
			case 'O':
				oWins++
			}
		}
	}
	// more then one win
	if xWins > 0 && oWins > 0 {
		return "", fmt.Errorf("both players cannot win")
	}
	// played on after win
	if xWins > 0 && x == o {
		return "", fmt.Errorf("O played after X won")
	}
	if oWins > 0 && x > o {
		return "", fmt.Errorf("X played after O won")
	}
	// one winner
	if xWins > 0 || oWins > 0 {
		return Win, nil
	}
	// draw
	if x+o == 9 {
		return Draw, nil
	}
	return Ongoing, nil
}
