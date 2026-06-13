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

func TestSession_Glyphs_CurrentRune(t *testing.T) {
	s := NewSession("cat")
	typeString(&s, "cx")
	g := s.Glyphs()

	if g[1].current != 'x' {
		t.Errorf("wrong glyph current = %q, want 'x'", g[1].current)
	}
	if g[2].current != 't' {
		t.Errorf("pending glyph current = %q, want 't'", g[2].current)
	}
}

func TestSession_Glyphs_Boundaries(t *testing.T) {
	s := NewSession("ab")
	typeString(&s, "abc")

	glyphs := s.Glyphs()
	if len(glyphs) != 2 {
		t.Fatalf("Glyphs() length = %d, want 2", len(glyphs))
	}

	last := glyphs[len(glyphs)-1]
	if last.state != Correct {
		t.Errorf("last glyph state = %v, want Correct", last.state)
	}
}

func TestSession_Type_HandlesSpecialRunes(t *testing.T) {
	targets := []struct {
		name   string
		target string
		want   int // glyph count
	}{
		{name: "spaces between words", target: "be water my friend", want: 18},
		{name: "punctuation", target: "Don't, stop!", want: 12},
		{name: "symbols and digits", target: "1 + 1 = 2 & 3", want: 13},
		{name: "accented latin", target: "café crème", want: 10},
		{name: "em dash and quotes", target: `“go” — now`, want: 10},
		{name: "tab inside text", target: "a\tb", want: 3},
	}

	for _, tt := range targets {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSession(tt.target)
			typeString(&s, tt.target)

			glyphs := s.Glyphs()
			if len(glyphs) != tt.want {
				t.Fatalf("glyph count = %d, want %d (rune count, not bytes)", len(glyphs), tt.want)
			}
			for i, g := range glyphs {
				if g.state != Correct {
					t.Errorf("glyph %d (%q) state = %v, want Correct", i, g.current, g.state)
				}
			}
		})
	}
}

func TestSession_Type_SpaceIsSymbol(t *testing.T) {
	s := NewSession("a b")
	typeString(&s, "axb")

	want := []CharState{Correct, Wrong, Correct}
	if got := states(s.Glyphs()); !slices.Equal(got, want) {
		t.Errorf("states = %v, want %v", got, want)
	}

	s2 := NewSession("abc")
	typeString(&s2, "a c")
	want2 := []CharState{Correct, Wrong, Correct}
	if got := states(s2.Glyphs()); !slices.Equal(got, want2) {
		t.Errorf("states = %v, want %v", got, want2)
	}
}

func TestSession_Backspace(t *testing.T) {
	t.Run("restores a wrong char to pending", func(t *testing.T) {
		s := NewSession("cat")
		typeString(&s, "cx")
		s.Backspace()

		want := []CharState{Correct, Pending, Pending}
		if got := states(s.Glyphs()); !slices.Equal(got, want) {
			t.Errorf("states = %v, want %v", got, want)
		}
	})

	t.Run("on an empty session does nothing", func(t *testing.T) {
		s := NewSession("cat")
		s.Backspace()

		want := []CharState{Pending, Pending, Pending}
		if got := states(s.Glyphs()); !slices.Equal(got, want) {
			t.Errorf("states = %v, want %v", got, want)
		}
	})

	t.Run("retyping after backspace tracks the cursor", func(t *testing.T) {
		s := NewSession("cat")
		typeString(&s, "cx")
		s.Backspace()
		s.Type('a')

		want := []CharState{Correct, Correct, Pending}
		if got := states(s.Glyphs()); !slices.Equal(got, want) {
			t.Errorf("states = %v, want %v", got, want)
		}
	})

	t.Run("backspace exposes the typed rune underneath", func(t *testing.T) {
		s := NewSession("cat")
		typeString(&s, "cat")
		s.Backspace()

		g := s.Glyphs()
		if g[2].state != Pending {
			t.Errorf("position 2 state = %v, want Pending", g[2].state)
		}
		if g[2].current != 't' {
			t.Errorf("position 2 current = %q, want 't'", g[2].current)
		}
	})
}
