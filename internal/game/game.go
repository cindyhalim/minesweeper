package game

import (
	"minesweeper/internal/components/board"
	"minesweeper/internal/components/help"
	"minesweeper/internal/components/keys"
	"minesweeper/internal/minesweeper"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)


type GameStatus int

const (
	IN_PROGRESS GameStatus = 0
	WIN GameStatus= 1
	LOSE GameStatus = 2
)
type Model struct {
	mode minesweeper.Mode
	board board.Model
	help help.Model
	status GameStatus
	highScore int
	width, height int
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd, boardCmd, helpCmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Quit):
			cmd = tea.Quit
		case key.Matches(msg, keys.Keys.Reset):
			m.resetGame()
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case board.GameWin:
		m.status = WIN
	case board.GameLose:
		m.status = LOSE
	}
	m.help, helpCmd = m.help.Update(msg)
	if m.status == IN_PROGRESS {
		m.board, boardCmd = m.board.Update(msg)
	}
	
	cmds = append(cmds, cmd, boardCmd, helpCmd)
	return m, tea.Batch(cmds...)
}

func (m *Model) resetGame() {
	m.board = board.NewModel(m.mode.Rows, m.mode.Cols, m.mode.Mines)
	m.status = IN_PROGRESS
}

func (m Model) View() string {
	boardView := lipgloss.JoinVertical(lipgloss.Center, m.board.View())
	statusView := formatGameStatus(m.status)
	mainView := lipgloss.JoinVertical(lipgloss.Center, statusView, boardView, m.help.View() )
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, mainView)
}

func NewModel(mode minesweeper.Mode) Model {
	return Model{
		mode: mode, 
		highScore: 0,
		board: board.NewModel(mode.Rows, mode.Cols, mode.Mines),
		help: help.NewModel(),
		status: IN_PROGRESS,
	}
}