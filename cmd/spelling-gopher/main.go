package main

import (
	"flag"
	"fmt"
	"log"

	tea "charm.land/bubbletea/v2"
	"github.com/Perchinka/SpellingGopher/internal/game"
	"github.com/Perchinka/SpellingGopher/internal/infra/csvquotes"
	"github.com/Perchinka/SpellingGopher/internal/quote"
	"github.com/Perchinka/SpellingGopher/internal/tui"
)

var version = "dev"

var (
	showVersion   = flag.Bool("version", false, "print version and exit")
	csvQuotesPath = flag.String("csv", "", "path to a custom quotes CSV (It has to have columns text,author)")
)

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Println(version)
		return
	}

	repo, err := newCsvRepo(*csvQuotesPath)
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

func newCsvRepo(path string) (*csvquotes.Repository, error) {
	if path != "" {
		return csvquotes.NewFromFile(path)
	}
	return csvquotes.New()
}
