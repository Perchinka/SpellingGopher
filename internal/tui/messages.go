package tui

import (
	"time"

	"github.com/Perchinka/SpellingGopher/internal/game"
	"github.com/Perchinka/SpellingGopher/internal/quote"
)

type quoteMsg struct{ quote quote.Quote }
type sessionFinishedMsg struct{ session *game.Session }
type restartGameMsg struct{}
type errMsg struct{ err error }
type tickMsg time.Time
