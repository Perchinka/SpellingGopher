package tui

import (
	"context"

	tea "charm.land/bubbletea/v2"
	"perchinka.github.io/spelling-gopher/internal/game"
)

func fetchQuote(quotes quoteSource) tea.Cmd {
	return func() tea.Msg {
		q, err := quotes.Random(context.Background())
		if err != nil {
			return errMsg{err}
		}
		return quoteMsg{q}
	}
}

func finished(s *game.Session) tea.Cmd {
	return func() tea.Msg { return sessionFinishedMsg{session: s} }
}

func restartGame() tea.Cmd {
	return func() tea.Msg { return restartGameMsg{} }
}
