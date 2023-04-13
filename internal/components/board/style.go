package board

import (
	"minesweeper/internal/minesweeper"
	"strconv"

	"github.com/charmbracelet/lipgloss"
)
const (
	HIDDEN_COLOR = "#dad8d4"
	FOCUS_COLOR = "#c2c0bc"
	BORDER_LIGHT_COLOR = "#f3f0ec"
	BORDER_DARK_COLOR = "#61605e"
)

var textToANSIMap = map[string]string {
	"1": "21",
	"2": "28",
	"3": "196",
	"4": "55",
	"5": "88",
	"6": "31",
	"7": "16",
	"8": "8",
	"X": "124",
}

var (
	cellStyle = func(value string, isInFocus bool) lipgloss.Style {
		style := lipgloss.NewStyle().Width(3).Bold(true).Border(lipgloss.ThickBorder()).Align(lipgloss.Center).BorderForeground(lipgloss.Color(BORDER_LIGHT_COLOR), lipgloss.Color(BORDER_DARK_COLOR), lipgloss.Color(BORDER_DARK_COLOR), lipgloss.Color(BORDER_LIGHT_COLOR)).ColorWhitespace(true)
		

		if _, ok := textToANSIMap[value]; ok {
			style = style.Foreground(lipgloss.Color(textToANSIMap[value]))
		}

		if isInFocus {
			style = style.Background(lipgloss.Color(FOCUS_COLOR)).BorderBackground(lipgloss.Color(FOCUS_COLOR))
		} else {
			style = style.Background(lipgloss.Color(HIDDEN_COLOR)).BorderBackground(lipgloss.Color(HIDDEN_COLOR))
		}

		return style
	}
	boardStyle = lipgloss.NewStyle().Margin(1)
)


func formatCell(cell Cell, isInFocus bool) string {
	styledValue := " "

	if cell.state != HIDDEN  {
		if cell.state == FLAGGED {
			styledValue = "X"
			return cellStyle(styledValue, isInFocus).Render(styledValue)
		} else if cell.value == minesweeper.MINE_CELL {
			styledValue = "💣"
		} else if cell.value != minesweeper.EMPTY_CELL {
			styledValue = strconv.Itoa(cell.value)
		}
	}

	return cellStyle(styledValue, isInFocus).Render(styledValue)
}

func makeInline(blocks string, block string) string {
	return lipgloss.JoinHorizontal(lipgloss.Center, blocks, block)
}

func formatRow(row string) string {
	return lipgloss.NewStyle().Render(row) + "\n"
}
