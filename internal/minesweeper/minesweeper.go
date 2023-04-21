package minesweeper

import (
	"math/rand"
	"time"
)

type BoardCellState int

const (
	MINE_CELL int = -1	
	EMPTY_CELL int = 0
)

type Board [][]int 

func (b Board) getRandomCell() (int, int) {
	row := len(b)
	col := len(b[0])

	currTime := time.Now().UnixNano()
	source := rand.NewSource(currTime)
	mineRow := rand.New(source).Intn(row)
	mineCol := rand.New(source).Intn(col)

	return mineRow, mineCol
}

func (b Board) updateCellValue(currRow int, currCol int) {
	rowLen := len(b)
	colLen := len(b[0])

	if currRow >= 0 && currRow < rowLen && currCol >= 0 && currCol < colLen && b[currRow][currCol] != MINE_CELL {
		b[currRow][currCol] += 1
	}

}

func New(width int, height int, row int, col int, mines int) Board {
	board := make (Board, width)
	for i := range board {
		board[i] = make([]int, height)
	}

	for i := 0; i < mines; i++ {
		mineRow, mineCol := board.getRandomCell()

		// ensure first clicked cell and adjacent cells isn't a mine
		for mineRow == row && mineCol == col ||
		mineRow == row-1 && mineCol == col ||
		mineRow == row+1 && mineCol == col ||
		mineRow == row && mineCol == col-1 ||
		mineRow == row && mineCol == col+1 ||
		mineRow == row-1 && mineCol == col-1 ||
		mineRow == row-1 && mineCol == col+1 ||
		mineRow == row+1 && mineCol == col-1 ||
		mineRow == row+1 && mineCol == col+1 {
			mineRow, mineCol = board.getRandomCell()
		}

		// ensure generated cell doesn't already contain mine
		for board[mineRow][mineCol] == MINE_CELL {
			mineRow, mineCol = board.getRandomCell()
		}
		
		board[mineRow][mineCol] = MINE_CELL

		// update mine's neighboring cell
		board.updateCellValue(mineRow - 1, mineCol) // N
		board.updateCellValue(mineRow + 1, mineCol) // S
		board.updateCellValue(mineRow, mineCol-1) // W
		board.updateCellValue(mineRow, mineCol+1) // E
		board.updateCellValue(mineRow-1, mineCol+1) // NE
		board.updateCellValue(mineRow-1, mineCol-1) // NW
		board.updateCellValue(mineRow+1, mineCol+1) // SE
		board.updateCellValue(mineRow+1, mineCol-1) // SW
	}

	return board
}
