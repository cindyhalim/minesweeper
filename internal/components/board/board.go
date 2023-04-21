package board

import (
	"minesweeper/internal/components/keys"
	"minesweeper/internal/minesweeper"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type CellState int

const (
	HIDDEN CellState = 0
	SHOWN CellState = 1
	FLAGGED CellState = 2
	MINE_CLICKED CellState = 3
	INCORRECT_FLAG CellState = 4
)

type Cell struct {
	value int
	state CellState
}

type coordinate struct {
	row, col int
}

type GameEnd struct {}

type Model struct {
	mines int
	board [][]Cell
	KeyMap keys.KeyMap
	isFirstClick bool
	cursor coordinate
}

func NewModel(rows int, cols int, mines int) Model {
	board := make([][]Cell, rows)
	for i := range board {
		board[i] = make([]Cell, cols)
	}

	return Model{
		mines: mines,
		board: board,
		KeyMap: keys.Keys,
		isFirstClick: true,
	}
}

func (m *Model) cursorUp() {
	if m.cursor.row > 0 {
		m.cursor.row--	
	} else {
		m.cursor.row = 0
	}
}


func (m *Model) cursorDown() {
	if m.cursor.row < len(m.board) {
		m.cursor.row++
	} else {
		m.cursor.row = len(m.board) - 1
	}
}

func (m *Model) cursorLeft() {
	if m.cursor.col > 0 {
		m.cursor.col--
	}  else {
		m.cursor.col = 0
	}
}

func (m *Model) cursorRight() {
	if m.cursor.col < len(m.board[0]) {
		m.cursor.col++
	} else {
		m.cursor.col = len(m.board[0]) - 1
	}
}

func (m *Model) revealCells() {
	coordinate := m.cursor

	if m.board[coordinate.row][coordinate.col].state == HIDDEN {
		currentCell := m.board[coordinate.row][coordinate.col]
		
		if currentCell.value == minesweeper.MINE_CELL {
			m.revealMines()
			m.endGame()
			return
		} 
		
		if currentCell.value == minesweeper.EMPTY_CELL{
			m.revealEmptyCells()
		}

		m.board[coordinate.row][coordinate.col].state = SHOWN
	}
}

func isInBound (board [][]Cell, row int, col int) bool {
	return row >= 0 && col >= 0 && row < len(board) && col < len(board[0])
}

func (m *Model) revealEmptyCells() {
	queue := []coordinate{}
	queue = append(queue, coordinate{row: m.cursor.row, col: m.cursor.col})

	for len(queue) > 0 {
		cellCoordinates := queue[0]
		queue = queue[1:]
		row := cellCoordinates.row
		col := cellCoordinates.col
		

		if isInBound(m.board, row, col) {
			if m.board[row][col].value == minesweeper.EMPTY_CELL && m.board[row][col].state == HIDDEN {
				queue = append(queue,
					coordinate{row: row-1, col: col},
					coordinate{row: row+1, col: col},
					coordinate{row: row, col: col-1},
					coordinate{row: row, col: col+1},
					coordinate{row: row-1, col: col-1},
					coordinate{row: row-1, col: col+1},
					coordinate{row: row+1, col: col-1},
					coordinate{row: row+1, col: col+1},
				)
			}
			m.board[row][col].state = SHOWN
		}
	}
}

func (m *Model) revealMines() {
	for i := range m.board {
		for j := range m.board[i] {
			if m.board[i][j].value != minesweeper.MINE_CELL && m.board[i][j].state == FLAGGED {
				m.board[i][j].state = INCORRECT_FLAG
			} else if m.board[i][j].value == minesweeper.MINE_CELL && m.board[i][j].state != FLAGGED {
				m.board[i][j].state = SHOWN
			}
		}
	}

	m.board[m.cursor.row][m.cursor.col].state = MINE_CLICKED
}

func (m *Model) endGame() tea.Cmd {
	return func() tea.Msg {
		return GameEnd{}
	}
}

func (m *Model) flag() {
	coordinate := m.cursor
	cellState := &m.board[coordinate.row][coordinate.col].state
	
	if *cellState == HIDDEN {
		*cellState = FLAGGED
	} else if *cellState == FLAGGED {
		*cellState = HIDDEN
	} 
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) getBoard() {
	width := len(m.board)
	height := len(m.board[0])
	minesweeper := minesweeper.New(width, height, m.cursor.row, m.cursor.col, m.mines)
	for i := range minesweeper {
		for j := range minesweeper[i] {
			m.board[i][j].value = minesweeper[i][j]
		}
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.Up):
			m.cursorUp()
		case key.Matches(msg, m.KeyMap.Down):
			m.cursorDown()
		case key.Matches(msg, m.KeyMap.Left):
			m.cursorLeft()
		case key.Matches(msg, m.KeyMap.Right):
			m.cursorRight()
		case key.Matches(msg, m.KeyMap.Click):
			if m.isFirstClick {
				m.getBoard()
				m.isFirstClick = false
			} 
			m.revealCells()
		case key.Matches(msg, m.KeyMap.Flag):
			m.flag()
		}
	}

	return m, nil
}

func (m Model) View() string {
	var board string
	for i := range m.board {
		row := ""
		for j := range m.board[i] {
			row = makeInline(row, formatCell(m.board[i][j], m.cursor.row == i && m.cursor.col == j))
		}
		board += formatRow(row)
	}

	return boardStyle.Render(board)
}

