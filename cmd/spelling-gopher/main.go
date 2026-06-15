package main

import (
	"fmt"
	"log"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/Perchinka/SpellingGopher/internal/game"
	"github.com/Perchinka/SpellingGopher/internal/infra/csvquotes"
	"github.com/Perchinka/SpellingGopher/internal/quote"
	"github.com/Perchinka/SpellingGopher/internal/tui"
)

var version = "dev"

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--version" {
		fmt.Println(version)
		return
	}

	repo, err := csvquotes.New()
	if err != nil {
		log.Fatal(err)
	}
	service := quote.NewService(repo)
	clock := game.RealClock{}

	model := tui.New(service, clock)

	if _, err := tea.NewProgram(model).Run(); err != nil {
		log.Fatal(err)
	}
}
