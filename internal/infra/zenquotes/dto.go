package zenquotes

import (
	"github.com/Perchinka/SpellingGopher/internal/quote"
)

type zenQuoteDTO struct {
	Quote  string `json:"q"`
	Author string `json:"a"`
}

func toDomain(dto zenQuoteDTO) quote.Quote {
	return quote.Quote{
		Text:   dto.Quote,
		Author: dto.Author,
	}
}
