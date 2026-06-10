package tui

import (
	"context"

	tea "charm.land/bubbletea/v2"
	"perchinka.github.io/spelling-gopher/internal/domain/quote"
)

func fetchQuote(service *quote.Service) tea.Cmd {
	return func() tea.Msg {
		q, err := service.Random(context.Background())
		if err != nil {
			return errMsg{err}
		}
		return quoteMsg{q}
	}
}
