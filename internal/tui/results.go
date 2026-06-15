package tui

import (
	"fmt"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
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
	if m.session == nil {
		return "no results yet"
	}
	st := m.session.Stats()

	body := fmt.Sprintf("%.0f wpm\n%.0f%% accuracy\n%s",
		st.WPM, st.Accuracy*100, st.Elapsed.Round(time.Second))

	content := lipgloss.NewStyle().
		Width(m.width).Height(m.height-1).
		Align(lipgloss.Center, lipgloss.Center).
		Render(body)

	footer := lipgloss.NewStyle().
		Align(lipgloss.Center).Width(m.width).
		Render("enter play again * esc quit")

	return lipgloss.JoinVertical(lipgloss.Left, content, footer)
}
