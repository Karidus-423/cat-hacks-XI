package user

import (
	"apothecary-journal/data"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)

	m.textinput, cmd = m.textinput.Update(msg)
	cmds = append(cmds, cmd)

	m.textinput, cmd = m.textinput.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		switch m.state {
		case MainView:
			switch key {
			case "q":
				return m, tea.Quit
			case "up", "k":
				if m.menuindex > 0 {
					m.menuindex--
				}
			case "down", "j":
				if m.menuindex < len(m.menuopts)-1 {
					m.menuindex++
				}
			case "enter", " ":
				_, ok := m.menuselect[m.menuindex]
				if ok {
					switch m.menuopts[m.menuindex] {
					case "Discover":
						m.state = DiscView
					case "Compendium":
						m.state = CompView
					}
				}
			}

		case DiscView:
		case CompView:
			switch key {
			case "q":
				return m, tea.Quit
			case "n":
				m.textinput.SetValue("")
				m.textinput.Focus()
				m.currflower = data.Flower{}
				m.state = CompView
			case "enter", " ":
				m.currflower = m.flowers[m.flowerindex]
				m.textarea.SetValue(m.currflower.Attributes)
				m.textarea.Focus()
				m.textarea.CursorEnd()
				m.currflower = m.flowers[m.flowerindex]
				m.state = FlowerView
			case "up", "k":
				if m.flowerindex > 0 {
					m.flowerindex--
				}
			case "down", "j":
				if m.flowerindex < len(m.flowers)-1 {
					m.flowerindex++
				}
			}
		}

	}
	return m, tea.Batch(cmds...)
}
