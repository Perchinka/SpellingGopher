package main

import (
	"log"

	tea "charm.land/bubbletea/v2"
	"perchinka.github.io/spelling-gopher/internal/infra/csvquotes"
	"perchinka.github.io/spelling-gopher/internal/quote"
	"perchinka.github.io/spelling-gopher/internal/tui"
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

	model := tui.New(service)

	if _, err := tea.NewProgram(model).Run(); err != nil {
		log.Fatal(err)
	}
}
