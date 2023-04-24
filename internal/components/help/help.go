package help

import (
	"minesweeper/internal/components/keys"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	help help.Model
	keys keys.KeyMap
}

func NewModel() Model {
	return Model {
		help: help.New(),
		keys: keys.Keys,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		}
	}
	
return m, nil
}

func (m Model) View() string {
	return lipgloss.NewStyle().MarginTop(1).MarginBottom(1).Render(m.help.View(m.keys))
}