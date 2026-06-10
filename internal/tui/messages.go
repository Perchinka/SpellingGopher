package tui

import "perchinka.github.io/spelling-gopher/internal/domain/quote"

type quoteMsg struct{ quote quote.Quote }
type errMsg struct{ err error }
