package board

import (
	"minesweeper/internal/minesweeper"
	"strconv"

	"github.com/charmbracelet/lipgloss"
)
const (
	HIDDEN_COLOR = lipgloss.Color("#dad8d4")
	FOCUS_COLOR = lipgloss.Color("#c2c0bc")
	SHOWN_COLOR = lipgloss.Color("#adadac")
	ERROR_COLOR = lipgloss.Color("#eb4034")
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
	cellStyle = func(value string, cellState CellState, isInFocus bool) lipgloss.Style {
		style := lipgloss.NewStyle().Width(3).Bold(true).Border(lipgloss.HiddenBorder()).Align(lipgloss.Center).ColorWhitespace(true)
		

		if _, ok := textToANSIMap[value]; ok {
			style = style.Foreground(lipgloss.Color(textToANSIMap[value]))
		}


		if isInFocus {
			style = style.Background(FOCUS_COLOR).BorderBackground(FOCUS_COLOR)
		} else if cellState == SHOWN {
			style = style.Background(SHOWN_COLOR).BorderBackground(SHOWN_COLOR)
		} else {
			style = style.Background(HIDDEN_COLOR).BorderBackground(HIDDEN_COLOR)
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
			return cellStyle(styledValue, cell.state, isInFocus).Render(styledValue)
		} else if cell.value == minesweeper.MINE_CELL {
			styledValue = "💣"
			return cellStyle(styledValue, cell.state, isInFocus).Background(ERROR_COLOR).BorderBackground(ERROR_COLOR).Render(styledValue)
		} else if cell.value != minesweeper.EMPTY_CELL {
			styledValue = strconv.Itoa(cell.value)
		}
	}

	return cellStyle(styledValue, cell.state, isInFocus).Render(styledValue)
}

func makeInline(blocks string, block string) string {
	return lipgloss.JoinHorizontal(lipgloss.Center, blocks, block)
}

func formatRow(row string) string {
	return lipgloss.NewStyle().Render(row) + "\n"
}
