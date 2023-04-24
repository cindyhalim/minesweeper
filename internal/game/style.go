package game

import "github.com/charmbracelet/lipgloss"

const (
	LOSE_COLOR = lipgloss.Color("#eb4034")
	WIN_COLOR = lipgloss.Color("#69d681")
)

var (
	statusStyle = lipgloss.NewStyle().Bold(true).MarginBottom(1)
	winStyle = func() lipgloss.Style {
		return statusStyle.Foreground(WIN_COLOR)
	}
	loseStyle = func() lipgloss.Style {
		return statusStyle.Foreground(LOSE_COLOR)
	}
)

func formatGameStatus(status GameStatus) string {
	if status == WIN {
		return winStyle().Render("you win!")
	}

	if status == LOSE {
		return loseStyle().Render("you lose!")
	}

	return statusStyle.Render(" ")
}