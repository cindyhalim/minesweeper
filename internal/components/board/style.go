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
	INCORRECT_FLAG_COLOR = lipgloss.Color("#fa857d")
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
	baseCellStyle = lipgloss.NewStyle().Width(3).Bold(true).Border(lipgloss.HiddenBorder()).Align(lipgloss.Center).ColorWhitespace(true)
	hiddenCellStyle = func(isInFocus bool) lipgloss.Style {
		if isInFocus {
			return baseCellStyle.Background(FOCUS_COLOR).BorderBackground(FOCUS_COLOR)
		} else {
			return baseCellStyle.Background(HIDDEN_COLOR).BorderBackground(HIDDEN_COLOR)
		}
	} 
	shownCellStyle = func(value string, isInFocus bool) lipgloss.Style {
		style := baseCellStyle

		if _, ok := textToANSIMap[value]; ok {
			style = style.Foreground(lipgloss.Color(textToANSIMap[value]))
		}


		if isInFocus {
			return style.Background(FOCUS_COLOR).BorderBackground(FOCUS_COLOR)
		} else {
			return style.Background(SHOWN_COLOR).BorderBackground(SHOWN_COLOR)
		}
	}
	mineClickedCellStyle = func() lipgloss.Style {
		return baseCellStyle.Background(ERROR_COLOR).BorderBackground(ERROR_COLOR)
	}
	incorrectFlagCellStyle = func() lipgloss.Style {
		return baseCellStyle.Background(INCORRECT_FLAG_COLOR).BorderBackground(INCORRECT_FLAG_COLOR)
	}
	boardStyle = lipgloss.NewStyle().Margin(1)
)

func getValue (value int, state CellState) string {
	if state == FLAGGED {
		return "X"
	}
	if value == minesweeper.MINE_CELL {
		return "💣"
	}
	if value == minesweeper.EMPTY_CELL {
		return " "
	}
	return strconv.Itoa(value)
}

func formatCell(cell Cell, isInFocus bool) string {

	if cell.state == HIDDEN {
		hiddenValue := " "
		return hiddenCellStyle(isInFocus).Render(hiddenValue)
	}

	formattedValue := getValue(cell.value, cell.state)
	if cell.state == MINE_CLICKED {
		return mineClickedCellStyle().Render(formattedValue)
	}

	if cell.state == INCORRECT_FLAG {
		return incorrectFlagCellStyle().Render(formattedValue)
	}

	return shownCellStyle(formattedValue, isInFocus).Render(formattedValue)

}


func makeInline(blocks string, block string) string {
	return lipgloss.JoinHorizontal(lipgloss.Center, blocks, block)
}

func formatRow(row string) string {
	return lipgloss.NewStyle().Render(row) + "\n"
}
