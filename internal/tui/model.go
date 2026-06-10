package tui

import (
	tea "charm.land/bubbletea/v2"
	"perchinka.github.io/spelling-gopher/internal/domain/quote"
)

type Model struct {
	quoteService *quote.Service
	quote        quote.Quote
	loading      bool
	err          error
}

func New(quoteService *quote.Service) Model {
	return Model{
		quoteService: quoteService,
		quote:        quote.Quote{Quote: "Nothing here yet", Author: "Me", CharacterCount: 0},
		loading:      false,
		err:          nil,
	}
}

func (m Model) Init() tea.Cmd {
	return fetchQuote(m.quoteService)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "r":
			m.loading = true
			return m, fetchQuote(m.quoteService)
		}
	case quoteMsg:
		m.quote, m.loading, m.err = msg.quote, false, nil
	case errMsg:
		m.loading, m.err = false, msg.err
	}
	return m, nil
}

func (m Model) View() tea.View {
	return tea.NewView(m.quote.Quote)
}
