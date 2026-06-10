package zenquotes

import "testing"

func TestToDomain(t *testing.T) {
	tests := []struct {
		name string
		dto  zenQuoteDTO
		want int
	}{
		{name: "ascii", dto: zenQuoteDTO{Quote: "be water", Author: "Lee"}, want: 8},
		{name: "unicode", dto: zenQuoteDTO{Quote: "café", Author: "x"}, want: 4},
		{name: "emoji", dto: zenQuoteDTO{Quote: "👍", Author: "x"}, want: 1},
		{name: "empty", dto: zenQuoteDTO{Quote: "", Author: "x"}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toDomain(tt.dto)

			if got.Quote != tt.dto.Quote {
				t.Errorf("Quote = %q, want %q", got.Quote, tt.dto.Quote)
			}
			if got.Author != tt.dto.Author {
				t.Errorf("Author = %q, want %q", got.Author, tt.dto.Author)
			}
			if got.CharacterCount != tt.want {
				t.Errorf("CharacterCount = %d, want %d", got.CharacterCount, tt.want)
			}
		})
	}
}
