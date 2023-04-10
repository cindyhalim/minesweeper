package main

import (
	"math/rand"
	"time"
)

type BoardCellState int

const (
	HIDDEN BoardCellState = 0
	SHOWN BoardCellState = 1
	FLAGGED BoardCellState = 2
)

const (
	MINE int = -1
)

type BoardCell struct {
	value int
	state BoardCellState
}

type Board [][]BoardCell


func (b Board) generateRandomCell() (int, int) {
	row := len(b)
	col := len(b[0])
	currTime := time.Now().UnixNano()
	source := rand.NewSource(currTime)
	mineRow := rand.New(source).Intn(row)
	mineCol := rand.New(source).Intn(col)

	return mineRow, mineCol
}

func (b Board) shouldUpdateCellValue(currRow int, currCol int) bool {
	rowLen := len(b)
	colLen := len(b[0])

	return currRow >= 0 && currRow < rowLen && currCol >= 0 && currCol < colLen && b[currRow][currCol].value != MINE
}

func createBoard(row int, col int, mines int) Board  {
	// create board based on dimensions
	board := make (Board, row)
	for i := range board {
		board[i] = make([]BoardCell, col)
	}
	
	// generate mines in random locations across board
	for i:=0; i < mines; i++ {
		mineRow, mineCol := board.generateRandomCell()

		// ensure generated cell doesn't already contain mine
		for board[mineRow][mineCol].value == MINE {
			mineRow, mineCol = board.generateRandomCell()
		}
		
		board[mineRow][mineCol].value = MINE

		// update mine's neighboring cell

		// N
		if board.shouldUpdateCellValue(mineRow - 1, mineCol) {
			board[mineRow-1][mineCol].value += 1
		}
		
		// S
		if board.shouldUpdateCellValue(mineRow + 1, mineCol) {
			board[mineRow+1][mineCol].value += 1
		}

		// W
		if board.shouldUpdateCellValue(mineRow, mineCol-1) {
			board[mineRow][mineCol-1].value += 1
		}

		// E
		if board.shouldUpdateCellValue(mineRow, mineCol+1) {
			board[mineRow][mineCol+1].value += 1
		}
		
		// NE
		if board.shouldUpdateCellValue(mineRow-1, mineCol+1) {
			board[mineRow-1][mineCol+1].value += 1
		}

		// NW
		if board.shouldUpdateCellValue(mineRow-1, mineCol-1) {
			board[mineRow-1][mineCol-1].value += 1
		}

		// SE
		if board.shouldUpdateCellValue(mineRow+1, mineCol+1) {
			board[mineRow+1][mineCol+1].value += 1
		}

		// SW
		if board.shouldUpdateCellValue(mineRow+1, mineCol-1) {
			board[mineRow+1][mineCol-1].value += 1
		}
	}

	return board
}

func main() {
	createBoard(9, 9, 10)

}
