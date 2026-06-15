package tui

import (
	"perchinka.github.io/spelling-gopher/internal/game"
	"perchinka.github.io/spelling-gopher/internal/quote"
)

type quoteMsg struct{ quote quote.Quote }
type sessionFinishedMsg struct{ session *game.Session }
type restartGameMsg struct{}
type errMsg struct{ err error }
