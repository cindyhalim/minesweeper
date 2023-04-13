package game

import (
	"minesweeper/internal/board"
	"minesweeper/internal/keys"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	board board.Model
	highScore int
	flagsRemaining int
}


func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Quit):
			cmd = tea.Quit
		}
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return "HELLO THIS IS GAME! \n" + m.board.View()
}

func NewModel(row int, col int, mines int) Model {
	return Model{
		board: board.NewModel(row, col, mines),
		highScore: 0,
		flagsRemaining: mines,
	}
}