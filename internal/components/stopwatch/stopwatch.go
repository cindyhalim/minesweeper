package stopwatch

import (
	"time"

	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	stopwatch stopwatch.Model
}

func NewModel() Model {
	return Model{
		stopwatch: stopwatch.NewWithInterval(time.Second),
	}
}

func (m *Model) Init() tea.Cmd {
	return m.stopwatch.Init()
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.stopwatch, cmd = m.stopwatch.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return lipgloss.NewStyle().Render("time elapsed: " + m.stopwatch.View())
}