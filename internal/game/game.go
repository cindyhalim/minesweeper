package game

import (
	"minesweeper/internal/components/board"
	"minesweeper/internal/components/keys"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	board board.Model
	highScore int
	width, height int
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd, boardCmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Quit):
			cmd = tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	m.board, boardCmd = m.board.Update(msg)
	cmds = append(cmds, cmd, boardCmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	mainView := lipgloss.JoinVertical(lipgloss.Center, m.board.View())
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, mainView)
}

func NewModel(row int, col int, mines int) Model {
	return Model{
		board: board.NewModel(row, col, mines),
		highScore: 0,
	}
}