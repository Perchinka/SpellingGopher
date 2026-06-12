package tui

import (
	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"perchinka.github.io/spelling-gopher/internal/domain/quote"
)

type Model struct {
	quoteService *quote.Service
	quote        quote.Quote
	loading      bool
	err          error

	spinner spinner.Model
}

func New(quoteService *quote.Service) Model {
	s := spinner.New(spinner.WithSpinner(spinner.Globe))
	return Model{
		quoteService: quoteService,
		quote:        quote.Quote{Text: "Nothing here yet", Author: "Me", CharacterCount: 0},
		loading:      true,
		err:          nil,
		spinner:      s,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(fetchQuote(m.quoteService), m.spinner.Tick)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "r":
			m.loading = true
			return m, tea.Batch(fetchQuote(m.quoteService), m.spinner.Tick)
		}
	case spinner.TickMsg:
		if !m.loading {
			return m, nil
		}
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case quoteMsg:
		m.quote, m.loading, m.err = msg.quote, false, nil
	case errMsg:
		m.loading, m.err = false, msg.err
	}
	return m, nil
}

func (m Model) View() tea.View {
	if m.loading {
		return tea.NewView(m.spinner.View() + " loading...")
	}
	if m.err != nil {
		return tea.NewView("error: " + m.err.Error())
	}
	return tea.NewView(m.quote.Text + "\n\t\t---" + m.quote.Author)
}
