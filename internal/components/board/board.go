package board

import (
	"minesweeper/internal/components/keys"
	"minesweeper/internal/minesweeper"

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

type Model struct {
	board [][]Cell
	KeyMap keys.KeyMap
}

func NewModel(row int, col int, mines int) Model {
	minesweeper := minesweeper.New(row, col, mines)
	board := make([][]Cell, row)
	for i := range board {
		board[i] = make([]Cell, col)
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

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	var board string
	for i := range m.board {
		row := ""
		for j := range m.board[i] {
			row += formatCell(m.board[i][j].value)
		}
		board += formatRow(row)
	}

	return boardStyle.Render(board)
}

