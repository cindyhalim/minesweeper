package board

import (
	"minesweeper/internal/components/keys"
	"minesweeper/internal/minesweeper"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type coordinate struct {
	row, col int
}

type GameEnd struct {}

type Model struct {
	mines int
	board minesweeper.Board
	KeyMap keys.KeyMap
	isFirstClick bool
	cursor coordinate
}

func NewModel(rows int, cols int, mines int) Model {
	board := make(minesweeper.Board, rows)
	for i := range board {
		board[i] = make([]minesweeper.Cell, cols)
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

func (m *Model) endGame() tea.Cmd {
	return func() tea.Msg {
		return GameEnd{}
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
			m.board[i][j].Value = minesweeper[i][j].Value
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
			minesweeper.RevealCells(&m.board, m.cursor.row, m.cursor.col)
		case key.Matches(msg, m.KeyMap.Flag):
			minesweeper.ToggleFlag(&m.board, m.cursor.row, m.cursor.col)
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

