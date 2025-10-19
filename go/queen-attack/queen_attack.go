package queenattack

import (
	"fmt"
	"math"
)

const (
	a = iota
	b
	c
	d
	e
	f
	g
	h
	lo = 0
	hi = 7
)

func positionToCoords(pos string) (col, row int, err error) {
	if len(pos) != 2 {
		return 0, 0, fmt.Errorf("invalid input: %q", pos)
	}
	col = int(pos[0] - 'a')
	row = 8 - int(pos[1]-'0')

	oob := col < lo || col > hi || row < lo || row > hi
	if oob {
		return -1, -1, fmt.Errorf("out of bounds: %q", pos)
	}
	return col, row, nil
}

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	whCol, whRow, err := positionToCoords(whitePosition)
	if err != nil {
		return false, err
	}
	blCol, blRow, err := positionToCoords(blackPosition)
	if err != nil {
		return false, err
	}
	if whRow == blRow && whCol == blCol { // both on the same spot
		return false, fmt.Errorf("same square")
	}
	colDiff := math.Abs(float64(whCol) - float64(blCol))
	rowDiff := math.Abs(float64(whRow) - float64(blRow))

	if whRow == blRow || whCol == blCol || colDiff == rowDiff {
		return true, nil
	}

	return false, nil
}
