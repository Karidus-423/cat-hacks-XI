package user

import (
	"apothecary-journal/data"
	"log"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	MainView uint = iota
	DiscView
	CompView
	FlowerView
)

type model struct {
	state       uint
	menuopts    []string
	menuselect  map[int]struct{}
	menuindex   int
	collection  *data.Collection
	flowers     []data.Flower
	currflower  data.Flower
	flowerindex int
	textinput   textinput.Model
	textarea    textarea.Model
	// options  list.Model
	// cursor   int
	// selected map[int]struct{}
}

func NewModel(collection *data.Collection) model {
	flowers, err := collection.GetFlowers()
	if err != nil {
		log.Fatalf("Unable to get flowers: %v", err)
	}
	return model{
		state:      MainView,
		menuopts:   []string{"Discover", "Compendium"},
		menuselect: make(map[int]struct{}),
		collection: collection,
		flowers:    flowers,
		textinput:  textinput.New(),
		textarea:   textarea.New(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
