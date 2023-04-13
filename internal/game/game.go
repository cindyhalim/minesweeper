package game

import (
	"minesweeper/internal/components/board"
	"minesweeper/internal/components/keys"
	"minesweeper/internal/components/stopwatch"
	"strconv"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	board board.Model
	stopwatch stopwatch.Model
	highScore int
	flagsRemaining int
	width, height int
}


func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd, boardCmd, stopwatchCmd tea.Cmd

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

	m.stopwatch, stopwatchCmd = m.stopwatch.Update(msg)
	m.board, boardCmd = m.board.Update(msg)

	cmds = append(cmds, cmd, boardCmd, stopwatchCmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	headerItemStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Bold(true)
	
	minesRemaining := "💣 remaining: " + strconv.Itoa(m.flagsRemaining)

	headerView := headerItemStyle.Width(m.width).Align(lipgloss.Center).Render(minesRemaining+"      ", m.stopwatch.View())

	mainView := lipgloss.JoinVertical(lipgloss.Center, headerView, m.board.View())

	mainView = lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, mainView)

	return mainView
}

func NewModel(row int, col int, mines int) Model {
	return Model{
		board: board.NewModel(row, col, mines),
		stopwatch: stopwatch.NewModel(),
		highScore: 0,
		flagsRemaining: mines,
	}
}