package zenquotes

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	quote "perchinka.github.io/spelling-gopher/internal/quote"
)

type Repository struct {
	client  *http.Client
	baseURL string
}

var (
	ErrUnexpectedStatus = errors.New("zenquotes: unexpected status")
	ErrEmptyResponse    = errors.New("zenquotes: empty response")
	ErrInvalidResponse  = errors.New("zenquotes: invalid response")
)

var _ quote.Repository = (*Repository)(nil)

func New(client *http.Client) *Repository {
	return &Repository{
		client:  client,
		baseURL: "https://zenquotes.io/api",
	}
}

func (r *Repository) Random(ctx context.Context) (quote.Quote, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, r.baseURL+"/random", nil)
	if err != nil {
		return quote.Quote{}, fmt.Errorf("zenquotes: build request: %w", err)
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return quote.Quote{}, fmt.Errorf("zenquotes: do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return quote.Quote{}, fmt.Errorf("%w: %d", ErrUnexpectedStatus, resp.StatusCode)
	}

	var dtos []zenQuoteDTO
	if err := json.NewDecoder(resp.Body).Decode(&dtos); err != nil {
		return quote.Quote{}, fmt.Errorf("%w: %v", ErrInvalidResponse, err)
	}
	if len(dtos) == 0 {
		return quote.Quote{}, ErrEmptyResponse
	}

	return toDomain(dtos[0]), nil
}
