package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

// GOOS=windows GOARCH=amd64 go build -o app.exe
func main() {
	p := tea.NewProgram(newModel())
	// Always check errors even if they should not happen.
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
