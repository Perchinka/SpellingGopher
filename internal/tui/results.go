package tui

import (
	tea "charm.land/bubbletea/v2"
	"perchinka.github.io/spelling-gopher/internal/game"
	"perchinka.github.io/spelling-gopher/internal/quote"
)

type resultsModel struct {
	screenBase
	session *game.Session
	quote   quote.Quote
}

func newResults() resultsModel {
	return resultsModel{
		quote: quote.Quote{Text: "Nothing to see here", Author: "Perchinka"},
	}
}

func (m resultsModel) Init() tea.Cmd {
	return nil
}

func (m resultsModel) Update(msg tea.Msg) (resultsModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "enter":
			return m, restartGame()
		}
	}
	return m, nil
}

func (m resultsModel) View() string {
	return "Results screen"
}
