package main

import (
	"apothecary-journal/scanning"
	"apothecary-journal/user"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func main() {
	_, format, err := scanning.DecodeImage("./images/training/abutilon.png")
	if err != nil {
		fmt.Println("DecodeImage Failed")
		return
	}
	fmt.Printf("Decoded Image | Format: %s\n", format)
	p := tea.NewProgram(user.InitialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v\n", err)
		os.Exit(1)
	}
}
