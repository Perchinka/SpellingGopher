package csvquotes

import (
	"context"
	_ "embed"
	"encoding/csv"
	"io"
	"math/rand/v2"
	"os"
	"strings"

	"github.com/Perchinka/SpellingGopher/internal/quote"
)

var _ quote.Repository = (*Repository)(nil)

//go:embed quotes.csv
var quotesData string

type Repository struct {
	quotes []quote.Quote
}

func New() (*Repository, error) {
	return NewFromReader(strings.NewReader(quotesData))
}

func NewFromFile(path string) (*Repository, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return NewFromReader(f)
}

func NewFromReader(r io.Reader) (*Repository, error) {
	records, err := csv.NewReader(r).ReadAll()
	if err != nil {
		return nil, err
	}

	quotes := make([]quote.Quote, 0, len(records))
	for _, row := range records[1:] {
		quotes = append(quotes, quote.Quote{
			Text:   row[0],
			Author: row[1],
		})
	}
	return &Repository{
		quotes: quotes,
	}, nil
}

func (r *Repository) Random(ctx context.Context) (quote.Quote, error) {
	return r.quotes[rand.IntN(len(r.quotes))], nil
}
