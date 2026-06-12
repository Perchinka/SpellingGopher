package zenquotes

import (
	"unicode/utf8"

	"perchinka.github.io/spelling-gopher/internal/domain/quote"
)

type zenQuoteDTO struct {
	Quote  string `json:"q"`
	Author string `json:"a"`
}

func toDomain(dto zenQuoteDTO) quote.Quote {
	return quote.Quote{
		Text:           dto.Quote,
		Author:         dto.Author,
		CharacterCount: utf8.RuneCountInString(dto.Quote),
	}
}
