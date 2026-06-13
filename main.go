package main

import (
	"net/http"
	"time"

	tea "charm.land/bubbletea/v2"
	"perchinka.github.io/spelling-gopher/internal/infra/zenquotes"
	"perchinka.github.io/spelling-gopher/internal/quote"
	"perchinka.github.io/spelling-gopher/internal/tui"
)

func main() {
	client := &http.Client{Timeout: 5 * time.Second}
	repo := zenquotes.New(client)
	service := quote.NewService(repo)

	model := tui.New(service)

	if _, err := tea.NewProgram(model).Run(); err != nil {
	}
}
