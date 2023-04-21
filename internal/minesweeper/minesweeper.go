package minesweeper

import (
	"math/rand"
	"time"
)

type CellState int
type CellValue int
type Coordinate struct {
	row int
	col int
}
type Cell struct {
	Value CellValue
	State CellState
}

type Board [][]Cell 

const (
	MINE_CELL CellValue = -1	
	EMPTY_CELL CellValue = 0

	HIDDEN CellState = 0
	SHOWN CellState = 1
	FLAGGED CellState = 2
	MINE_CLICKED CellState = 3
	INCORRECT_FLAG CellState = 4
)

func isInBound (board [][]Cell, row int, col int) bool {
	return row >= 0 && col >= 0 && row < len(board) && col < len(board[0])
}

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

	if currRow >= 0 && currRow < rowLen && currCol >= 0 && currCol < colLen && b[currRow][currCol].Value != MINE_CELL {
		b[currRow][currCol].Value += 1
	}

}

func (board Board) revealEmptyCells( row int, col int) {
	queue := []Coordinate{}
	queue = append(queue, Coordinate{row, col})

	for len(queue) > 0 {
		cellCoordinates := queue[0]
		queue = queue[1:]
		row := cellCoordinates.row
		col := cellCoordinates.col
		

		if isInBound(board, row, col) {
			if board[row][col].Value == EMPTY_CELL && board[row][col].State == HIDDEN {
				queue = append(queue,
					Coordinate{row: row-1, col: col},
					Coordinate{row: row+1, col: col},
					Coordinate{row: row, col: col-1},
					Coordinate{row: row, col: col+1},
					Coordinate{row: row-1, col: col-1},
					Coordinate{row: row-1, col: col+1},
					Coordinate{row: row+1, col: col-1},
					Coordinate{row: row+1, col: col+1},
				)
			}
			board[row][col].State = SHOWN
		}
	}
}

func (board Board) revealMines(row int, col int) {
	for i := range board {
		for j := range board[i] {
			if board[i][j].Value != MINE_CELL && board[i][j].State == FLAGGED {
				board[i][j].State = INCORRECT_FLAG
			} else if board[i][j].Value == MINE_CELL && board[i][j].State != FLAGGED {
				board[i][j].State = SHOWN
			}
		}
	}

	board[row][col].State = MINE_CLICKED
}


func New(width int, height int, row int, col int, mines int) Board {
	board := make (Board, width)
	for i := range board {
		board[i] = make([]Cell, height)
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
		for board[mineRow][mineCol].Value == MINE_CELL {
			mineRow, mineCol = board.getRandomCell()
		}
		
		board[mineRow][mineCol].Value = MINE_CELL

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

func RevealCells(b *Board, row int, col int) {
	board := *b

	if board[row][col].State == HIDDEN {
		currentCell := board[row][col]
		
		if currentCell.Value == MINE_CELL {
			board.revealMines(row, col)
			return
		} 
		
		if currentCell.Value == EMPTY_CELL{
			board.revealEmptyCells(row, col)
		}

		board[row][col].State = SHOWN
	}
}

func ToggleFlag(b *Board, row int, col int) {
	board := *b

	cellState := &board[row][col].State
	
	if *cellState == HIDDEN {
		*cellState = FLAGGED
	} else if *cellState == FLAGGED {
		*cellState = HIDDEN
	} 
}
