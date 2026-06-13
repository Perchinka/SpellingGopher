package game

import (
	"slices"
	"testing"
)

func typeString(s *Session, text string) {
	for _, r := range text {
		s.Type(r)
	}
}

func states(glyphs []Glyph) []CharState {
	out := make([]CharState, len(glyphs))
	for i, g := range glyphs {
		out[i] = g.state
	}
	return out
}

func TestSession_Glyphs_States(t *testing.T) {
	tests := []struct {
		name   string
		target string
		typed  string
		want   []CharState
	}{
		{
			name:   "fresh session is all pending",
			target: "cat",
			typed:  "",
			want:   []CharState{Pending, Pending, Pending},
		},
		{
			name:   "correct then wrong then pending",
			target: "cat",
			typed:  "cx",
			want:   []CharState{Correct, Wrong, Pending},
		},
		{
			name:   "All correct",
			target: "cat",
			typed:  "cat",
			want:   []CharState{Correct, Correct, Correct},
		},
		{
			name:   "All wrong",
			target: "cat",
			typed:  "xfz",
			want:   []CharState{Wrong, Wrong, Wrong},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSession(tt.target)
			typeString(&s, tt.typed)
			if got := states(s.Glyphs()); !slices.Equal(got, tt.want) {
				t.Errorf("states = %v, want %v", got, tt.want)
			}
		})
	}
}
