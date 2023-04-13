package board

import (
	"minesweeper/internal/minesweeper"
	"strconv"

	"github.com/charmbracelet/lipgloss"
)
const (
	HIDDEN_COLOR = "#dad8d4"
	FOCUS_COLOR = "#c2c0bc"
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
}

var (
	cellStyle = func(value string, isInFocus bool) lipgloss.Style {
		style := lipgloss.NewStyle().Width(3).Bold(true).Border(lipgloss.HiddenBorder()).ColorWhitespace(true).Align(lipgloss.Center)

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

	rowStyle = lipgloss.NewStyle()
	boardStyle = lipgloss.NewStyle().Margin(1)
)


func formatCell(value int, isInFocus bool) string {
	styledValue := " "

	if value == minesweeper.MINE_CELL {
		styledValue = "💣"
	} else if value != minesweeper.EMPTY_CELL {
		styledValue = strconv.Itoa(value)
	}

	return cellStyle(styledValue, isInFocus).Render(styledValue)
}

func makeInline(blocks string, block string) string {
	return lipgloss.JoinHorizontal(lipgloss.Center, blocks, block)
}

func formatRow(row string) string {
	return rowStyle.Render(row) + "\n"
}
