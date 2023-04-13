package keys

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	Quit key.Binding
}

var Keys = KeyMap{
	Quit: key.NewBinding(key.WithKeys("q", "esc"), key.WithHelp("q/esc", "quit")),
}