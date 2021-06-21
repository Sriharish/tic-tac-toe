package ttt

import (
	"fmt"
)

const (
	// Board State Enumerations
	P1    = 'X'
	P2    = 'O'
	EMPTY = '-'

	// Game State Enumerations
	P1WIN      = "P1"
	P2WIN      = "P2"
	TIE        = "TIE"
	INPROGRESS = "INPROGRESS"

	maxTurns = 9
)

type Board struct {
	State [][]rune
}

type Game struct {
	P1Score     int
	P2Score     int
	CurrentTurn rune
	TurnCount   int
}

func CreateBoard(b *Board) {
	state := make([][]rune, 3)
	for i := range state {
		state[i] = []rune{EMPTY, EMPTY, EMPTY}
	}

	b.State = state
}

func InitGame(g *Game, firstTurn rune) (bool, error) {
	if firstTurn != P1 && firstTurn != P2 {
		return false, fmt.Errorf("[ttt]InitGame: Invalid player character. Please choose [%v] or [%v]\n", P1, P2)
	}
	g.P1Score = 0
	g.P2Score = 0
	g.CurrentTurn = firstTurn
	g.TurnCount = 0
	return true, nil
}

func PrintBoard(b *Board) {
	for _, val := range b.State {
		fmt.Println(val)
	}
	fmt.Println()
}

func EvalBoard(b *Board, gameState *Game, currentTurn rune) string {
	// Check rows
	for _, row := range b.State {
		for j, col := range row {
			if col != currentTurn {
				break
			} else if j == len(row)-1 {
				if currentTurn == P1 {
					return P1WIN
				}
				return P2WIN
			}
		}
	}

	// Check columns
	for i := 0; i < len(b.State); i++ { // row
		for j := 0; j < len(b.State[0]); j++ { // col
			if b.State[j][i] != currentTurn {
				break
			} else if j == len(b.State)-1 {
				if currentTurn == P1 {
					return P1WIN
				}
				return P2WIN
			}
		}
	}

	// Check diagonal from top left
	for i := 0; i < len(b.State); i++ {
		if b.State[i][i] != currentTurn {
			break
		} else if i == len(b.State)-1 {
			if currentTurn == P1 {
				return P1WIN
			}
			return P2WIN
		}
	}

	// Check diagonal from top right
	for i, j := 0, len(b.State)-1; i < len(b.State) && j >= 0; i, j = i+1, j-1 {
		if b.State[i][j] != currentTurn {
			break
		} else if i == len(b.State)-1 {
			if currentTurn == P1 {
				return P1WIN
			}
			return P2WIN
		}
	}

	// Check tie state
	if gameState.TurnCount >= maxTurns {
		return TIE
	}
	return INPROGRESS
}
