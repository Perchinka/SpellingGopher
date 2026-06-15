package tui

import (
	"context"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"perchinka.github.io/spelling-gopher/internal/game"
	"perchinka.github.io/spelling-gopher/internal/quote"
)

// for color codes https://github.com/fidian/ansi
// TODO don't forget to move this to the separate styles.go later
var (
	pendingStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	correctStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("10"))
	wrongStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	wrongSpaceStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Underline(true)
)

func styleFor(glyph game.Glyph) lipgloss.Style {
	switch glyph.State {
	case game.Correct:
		return correctStyle
	case game.Wrong:
		if glyph.IsSpace() {
			return wrongSpaceStyle
		}
		return wrongStyle
	default:
		return pendingStyle
	}
}

type quoteSource interface {
	Random(ctx context.Context) (quote.Quote, error)
}

type typingModel struct {
	screenBase
	quotes       quoteSource
	currentQuote quote.Quote
	loading      bool
	session      *game.Session
	clock        game.Clock
	err          error
}

func newTyping(quotes quoteSource, clock game.RealClock) typingModel {
	return typingModel{
		quotes:  quotes,
		loading: true,
		clock:   clock,
	}
}

func (m typingModel) Init() tea.Cmd {
	return tea.Batch(fetchQuote(m.quotes), tick())
}

func (m typingModel) Update(msg tea.Msg) (typingModel, tea.Cmd) {
	switch msg := msg.(type) {
	case quoteMsg:
		m.currentQuote, m.loading, m.err = msg.quote, false, nil
		m.session = game.NewSession(msg.quote.Text, m.clock)
	case errMsg:
		m.loading, m.err = false, msg.err
	case tea.KeyPressMsg:
		if m.session == nil {
			return m, nil
		}
		wasFinished := m.session.Finished()
		switch {
		case msg.Code == tea.KeyBackspace:
			m.session.Backspace()
		case msg.Text != "":
			for _, r := range msg.Text {
				m.session.Type(r)
			}
		}
		if !wasFinished && m.session.Finished() {
			return m, finished(m.session)
		}
	case tickMsg:
		if m.session == nil || m.session.Finished() {
			return m, nil
		}
		return m, tick()
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

	var b strings.Builder
	for _, glyph := range m.session.Glyphs() {
		b.WriteString(styleFor(glyph).Render(string(glyph.Current)))
	}
	return b.String()
}
