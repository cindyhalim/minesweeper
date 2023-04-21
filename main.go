package main

import (
	"fmt"
	"minesweeper/internal/game"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type BoardSettings struct {
	row int
	col int
	mines int
}

func main() {
	var difficulty BoardSettings
	var difficultyToBoardSettings = map[string]BoardSettings{
		"beginner": {row: 9, col: 9, mines: 10},
		"intermediate": {row: 16, col: 16, mines: 40},
		"expert": {row: 16, col: 30, mines: 99},
	}
	
	if len(os.Args) < 2 {
		difficulty = difficultyToBoardSettings["beginner"]
	} else {
		if _, ok := difficultyToBoardSettings[os.Args[1]]; !ok {
			fmt.Printf("Invalid difficulty: %s\n", os.Args[1])
			os.Exit(1)
		}
		difficulty = difficultyToBoardSettings[os.Args[1]]
	}
	
	p := tea.NewProgram(game.NewModel(difficulty.row, difficulty.col, difficulty.mines), tea.WithAltScreen(), tea.WithMouseCellMotion())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}                        
