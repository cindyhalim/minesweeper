package keys

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	Quit key.Binding
	Up key.Binding
	Down key.Binding
	Left key.Binding
	Right key.Binding
	Click key.Binding
	Flag key.Binding
	Reset key.Binding
}

var Keys = KeyMap{
	Quit: key.NewBinding(key.WithKeys("q", "esc"), key.WithHelp("q/esc", "quit")),
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "move left"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "move right"),
	),
	Click: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("space", "reveal cell"),
	),
	Flag: key.NewBinding(
		key.WithKeys("x"),
		key.WithHelp("x", "flag cell"),
	),
	Reset: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "reset board"),
	),
}