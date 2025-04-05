package main

import (
	"apothecary-journal/data"
	"apothecary-journal/user"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	collection := &data.Collection{}
	if err := collection.Init(); err != nil {
		log.Fatalf("Unable to init collection: %v", err)
	}
	m := user.NewModel(collection)

	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		log.Fatalf("Unable to run program due to %v", err)
	}
}
