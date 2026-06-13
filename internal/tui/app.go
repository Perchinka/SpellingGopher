package tui

import (
	tea "charm.land/bubbletea/v2"
)

type screenBase struct {
	width, height int
}

func (b *screenBase) SetSize(w, h int) { b.width, b.height = w, h }

type screenID int

const (
	screenTyping screenID = iota
	screenResults
)

type Model struct {
	active        screenID
	typing        typingModel
	results       resultsModel
	width, height int
}

func New() Model {
	return Model{
		active:  screenTyping,
		typing:  newTyping(),
		results: newResults(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.active {
	case screenTyping:
		m.typing, _ = m.typing.Update(msg)
	}

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		// NOTE for two screens it's alright, but it feels wrong...
		// NOTE I guess I will add some interface for screens later and will store them in the map
		m.height, m.width = msg.Height, msg.Width
		m.typing.SetSize(m.width, m.height)
		m.results.SetSize(m.width, m.height)
	}
	return m, nil
}

func (m Model) View() tea.View {
	var content string
	switch m.active {
	case screenTyping:
		content = m.typing.View()
	case screenResults:
		content = m.results.View()

	}
	return tea.NewView(content)
}
