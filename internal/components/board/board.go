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
)

type Cell struct {
	value int
	state CellState
}

type coordinate struct {
	row, col int
}
type Model struct {
	board [][]Cell
	KeyMap keys.KeyMap
	cursor coordinate
}

func NewModel(rows int, cols int, mines int) Model {
	minesweeper := minesweeper.New(rows, cols, mines)
	board := make([][]Cell, rows)
	for i := range board {
		board[i] = make([]Cell, cols)
	}

	for i := range minesweeper {
		for j := range minesweeper[i] {
			board[i][j].value = minesweeper[i][j]
		}
	}

	return Model{
		board: board,
		KeyMap: keys.Keys,
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

func (m Model) Init() tea.Cmd {
	return nil
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
		}
	}

	return m, nil
}

func (m Model) View() string {
	var board string
	for i := range m.board {
		row := ""
		for j := range m.board[i] {
			row = makeInline(row, formatCell(m.board[i][j].value, m.cursor.row == i && m.cursor.col == j))
		}
		board += formatRow(row)
	}

	return boardStyle.Render(board)
}

