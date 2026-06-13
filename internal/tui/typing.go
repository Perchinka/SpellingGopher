package tui

import (
	"context"

	tea "charm.land/bubbletea/v2"
	"perchinka.github.io/spelling-gopher/internal/quote"
)

type quoteSource interface {
	Random(ctx context.Context) (quote.Quote, error)
}

type typingModel struct {
	screenBase
	quotes       quoteSource
	currentQuote quote.Quote
	loading      bool
	err          error
}

func newTyping(quotes quoteSource) typingModel {
	return typingModel{
		quotes:  quotes,
		loading: true,
	}
}

func (m typingModel) Init() tea.Cmd {
	return fetchQuote(m.quotes)
}

func (m typingModel) Update(msg tea.Msg) (typingModel, tea.Cmd) {
	switch msg := msg.(type) {
	case quoteMsg:
		m.currentQuote, m.loading, m.err = msg.quote, false, nil
	case errMsg:
		m.loading, m.err = false, msg.err
	}
	return m, nil
}

func (m typingModel) View() string {
	if m.loading {
		return "loading..."
	}
	if m.err != nil {
		return "error: " + m.err.Error()
	}
	return m.currentQuote.Text + "\n\t\t-" + m.currentQuote.Author
}
