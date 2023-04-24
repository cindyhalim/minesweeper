package main

import (
	"fmt"
	"minesweeper/internal/game"
	"minesweeper/internal/minesweeper"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)


func main() {
	var mode minesweeper.Mode
	modesMap := map[string]minesweeper.Mode{
		"beginner":   minesweeper.BEGINNER,
		"intermediate": minesweeper.INTERMEDIATE,
		"expert":   minesweeper.EXPERT,
	}
	
	if len(os.Args) < 2 {
		mode = minesweeper.BEGINNER
	} else {
		if _, ok := modesMap[os.Args[1]]; !ok {
			fmt.Printf("Invalid mode: %s\n", os.Args[1])
			os.Exit(1)
		}
		mode = modesMap[os.Args[1]]
	}
	
	p := tea.NewProgram(game.NewModel(mode), tea.WithAltScreen(), tea.WithMouseCellMotion())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}                        
