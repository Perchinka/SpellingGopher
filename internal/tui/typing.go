package tui

import (
	tea "charm.land/bubbletea/v2"
	"perchinka.github.io/spelling-gopher/internal/quote"
)

type typingModel struct {
	screenBase
	quote quote.Quote
}

func newTyping() typingModel {
	return typingModel{
		quote: quote.Quote{Text: "Nothing to see here", Author: "Perchinka"},
	}
}

func (m typingModel) Init() tea.Cmd {
	return nil
}

func (m typingModel) Update(msg tea.Msg) (typingModel, tea.Cmd) {
	// switch msg := msg.(type) {
	// case tea.KeyPressMsg:
	// }
	return m, nil
}

func (m typingModel) View() tea.View {
	return tea.NewView("test")
}
