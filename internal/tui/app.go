package tui

import (
	tea "charm.land/bubbletea/v2"
	"perchinka.github.io/spelling-gopher/internal/game"
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

func New(quotes quoteSource, clock game.RealClock) Model {
	return Model{
		active:  screenTyping,
		typing:  newTyping(quotes, clock),
		results: newResults(),
	}
}

func (m Model) Init() tea.Cmd {
	return m.typing.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		// NOTE for two screens it's alright, but it feels wrong...
		// NOTE I guess I will add some interface for screens later and will store them in the map
		m.width, m.height = msg.Width, msg.Height
		m.typing.SetSize(m.width, m.height)
		m.results.SetSize(m.width, m.height)
		return m, nil
	case sessionFinishedMsg:
		m.active = screenResults
		m.results.session = msg.session
		return m, nil

	case restartGameMsg:
		m.active = screenTyping
		m.typing.loading = true
		m.typing.session = nil
		return m, m.typing.Init()
	}

	var cmd tea.Cmd
	switch m.active {
	case screenTyping:
		m.typing, cmd = m.typing.Update(msg)
	case screenResults:
		m.results, cmd = m.results.Update(msg)
	}

	return m, cmd
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
