package board

import (
	"minesweeper/internal/components/keys"
	"minesweeper/internal/components/stopwatch"
	"minesweeper/internal/minesweeper"
	"strconv"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type coordinate struct {
	row, col int
}

type GameWin struct {}
type GameLose struct {}

type Model struct {
	totalMines int
	flagsRemaining int
	board minesweeper.Board
	stopwatch stopwatch.Model
	KeyMap keys.KeyMap
	isFirstClick bool
	cursor coordinate
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

func createEmptyBoard(rows int, cols int) minesweeper.Board {
	board := make(minesweeper.Board, rows)
	for i := range board {
		board[i] = make([]minesweeper.Cell, cols)
	}

	return board
}

func (m *Model) fillBoard() {
	width := len(m.board)
	height := len(m.board[0])
	minesweeper := minesweeper.New(width, height, m.cursor.row, m.cursor.col, m.totalMines)
	for i := range minesweeper {
		for j := range minesweeper[i] {
			m.board[i][j].Value = minesweeper[i][j].Value
		}
	}
}

func (m *Model) countFlagsRemaining() {
	if m.board[m.cursor.row][m.cursor.col].State == minesweeper.FLAGGED {
		m.flagsRemaining--
	} else if m.board[m.cursor.row][m.cursor.col].State == minesweeper.HIDDEN {
		m.flagsRemaining++
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) gameLose() tea.Cmd {
	return func() tea.Msg {
		return GameLose{}
	}
}

func (m *Model) gameWin() tea.Cmd {
	return func() tea.Msg {
		return GameWin{}
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd, stopwatchCmd tea.Cmd

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
				m.fillBoard()
				m.isFirstClick = false
				cmd = m.stopwatch.Init()
			} 
			minesweeper.RevealCells(&m.board, m.cursor.row, m.cursor.col)

			if m.board[m.cursor.row][m.cursor.col].Value == minesweeper.MINE_CELL {
				return m, m.gameLose()
			}
	
			if m.flagsRemaining == 0 {
				if m.checkBoard() {
					return m, m.gameWin()
				}
			}
			
		case key.Matches(msg, m.KeyMap.Flag):
			minesweeper.ToggleFlag(&m.board, m.cursor.row, m.cursor.col)
			m.countFlagsRemaining()
		}
	}
	m.stopwatch, stopwatchCmd = m.stopwatch.Update(msg)
	cmds = append(cmds, cmd, stopwatchCmd)

	return m, tea.Batch(cmds...)
}

func (m Model) checkBoard() bool {
	for i := range m.board {
		for j := range m.board[i] {
			cell := m.board[i][j]

			if cell.Value == minesweeper.EMPTY_CELL && cell.State != minesweeper.SHOWN {
				return false
			}

			if cell.Value == minesweeper.MINE_CELL && cell.State != minesweeper.FLAGGED {
				return false
			}
		}
	}

	return true
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
	
	flagsRemainingText := "🚩 remaining: " + strconv.Itoa(m.flagsRemaining)
	headerView := headerStyle.Render(flagsRemainingText+"      ", m.stopwatch.View())
	mainView := lipgloss.JoinVertical(lipgloss.Center, headerView, boardStyle.Render(board))

	return mainView
}


func NewModel(rows int, cols int, mines int) Model {
	board := createEmptyBoard(rows, cols)
	
	return Model{
		totalMines: mines,
		flagsRemaining: mines,
		board: board,
		stopwatch: stopwatch.NewModel(),
		KeyMap: keys.Keys,
		isFirstClick: true,
	}
}

