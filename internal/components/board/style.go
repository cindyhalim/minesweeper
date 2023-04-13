package board

import (
	"minesweeper/internal/minesweeper"
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

var (
	cellStyle = func(value string) lipgloss.Style {
		style := lipgloss.NewStyle().Width(5).ColorWhitespace(true).Align(lipgloss.Center).Bold(true).Background(lipgloss.Color("8"))

		if _, ok := textToANSIMap[value]; ok {
			style = style.Foreground(lipgloss.Color(textToANSIMap[value]))
		}

		return style
	}

	rowStyle = func() lipgloss.Style {
		return lipgloss.NewStyle().PaddingBottom(1).PaddingTop(1).Align(lipgloss.Center).Background(lipgloss.Color("8"))
	}

	boardStyle = lipgloss.NewStyle().Margin(1)
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


func formatCell(value int) string {
	styledValue := " "

	if value == minesweeper.MINE_CELL {
		styledValue = "💣"
	} else if value != minesweeper.EMPTY_CELL {
		styledValue = strconv.Itoa(value)
	}

	return cellStyle(styledValue).Render(styledValue)
}

func formatRow(row string) string {
	return rowStyle().Render(row) + "\n"
}
