package user

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	appNameStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("99")).
			Padding(0, 1)

	faintStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("255")).
			Faint(true)

	enumeratorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("99")).MarginRight(1)
)

func (m model) View() string {
	s := appNameStyle.Render("Apothecary Journal") + "\n\n"
	if m.state == MainView {
		for i, chosen := range m.menuopts {
			prefix := ""
			if i == m.menuindex {
				prefix = ">"
			}

			s += fmt.Sprintf("%s | %s", prefix, chosen)
		}
		s += "n: New Flower, q - quit"
	} else if m.state == CompView {
		for i, n := range m.flowers {
			prefix := " "
			if i == m.flowerindex {
				prefix = ">"
			}

			shortattr := strings.ReplaceAll(n.Attributes, "\n", " ")
			if len(shortattr) > 30 {
				shortattr = shortattr[:30]
			}

			s += enumeratorStyle.Render(prefix) +
				n.Title + " | " +
				faintStyle.Render(shortattr) + "\n\n"
		}
		s += faintStyle.Render("n: New Flower, q - quit")
	}

	return s
}
