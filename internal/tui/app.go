package tui

import (
	tea "charm.land/bubbletea/v2"
)

type screenID int

const (
	screenTyping screenID = iota
	screenResults
)

type Model struct {
	active  screenID
	typing  typingModel
	results resultsModel
	err     error
}

func New() Model {
	return Model{
		active:  screenTyping,
		typing:  newTyping(),
		results: newResults(),
		err:     nil,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() tea.View {
	return tea.NewView("test")
}
