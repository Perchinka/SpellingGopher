package game

import "testing"

func TestGlyph_IsSpace(t *testing.T) {
	tests := []struct {
		name    string
		current rune
		want    bool
	}{
		{name: "literal space is a space", current: ' ', want: true},
		{name: "letter is not a space", current: 'a', want: false},
		{name: "tab is not a space", current: '\t', want: false},
		{name: "zero rune is not a space", current: 0, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Glyph{Current: tt.current}
			if got := g.IsSpace(); got != tt.want {
				t.Errorf("IsSpace(%q) = %v, want %v", tt.current, got, tt.want)
			}
		})
	}
}
