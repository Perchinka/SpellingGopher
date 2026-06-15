package main

import (
	"log"

	tea "charm.land/bubbletea/v2"
	"github.com/Perchinka/SpellingGopher/internal/game"
	"github.com/Perchinka/SpellingGopher/internal/infra/csvquotes"
	"github.com/Perchinka/SpellingGopher/internal/quote"
	"github.com/Perchinka/SpellingGopher/internal/tui"
)

func main() {
	// client := &http.Client{Timeout: 15 * time.Second}
	// repo := zenquotes.New(client)
	repo, err := csvquotes.New()
	if err != nil {
		log.Fatal(err)
		return
	}
	service := quote.NewService(repo)
	clock := game.RealClock{}

	model := tui.New(service, clock)

	if _, err := tea.NewProgram(model).Run(); err != nil {
		log.Fatal(err)
	}
}
