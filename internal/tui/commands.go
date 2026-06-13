package tui

import (
	"context"

	tea "charm.land/bubbletea/v2"
)

// NOTE Pull it out like that to not block Update thread
func fetchQuote(quotes quoteSource) tea.Cmd {
	return func() tea.Msg {
		q, err := quotes.Random(context.Background())
		if err != nil {
			return errMsg{err}
		}
		return quoteMsg{q}
	}
}
